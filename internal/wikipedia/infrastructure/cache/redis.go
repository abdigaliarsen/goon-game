package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"goon-game/internal/wikipedia/config"
	"goon-game/pkg/utils"
	"time"
)

type redisCache struct {
	cfg    *config.Config
	logger utils.Logger
	rdb    *redis.Client
}

type RedisCacheIn struct {
	fx.In
	Cfg    *config.Config
	Logger utils.Logger
}

func New(in RedisCacheIn) Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     in.Cfg.RedisConfig.Addr,
		Password: in.Cfg.RedisConfig.Password,
		DB:       in.Cfg.RedisConfig.DB,
	})

	return &redisCache{
		cfg:    in.Cfg,
		logger: in.Logger,
		rdb:    rdb,
	}
}

func (r *redisCache) SetS(key, value string) error {
	st := r.rdb.Set(context.TODO(), key, value, 0)
	if st.Err() != nil {
		return st.Err()
	}

	return nil
}

func (r *redisCache) GetS(key string) (string, error) {
	val := r.rdb.Get(context.TODO(), key)
	if val.Err() != nil && !errors.Is(val.Err(), redis.Nil) {
		return "", val.Err()
	}

	return val.Val(), nil
}

func (r *redisCache) AddS(key, value string) error {
	timestamp := time.Now().Unix()

	_, err := r.rdb.ZAdd(context.TODO(), key, redis.Z{
		Score:  float64(timestamp),
		Member: value,
	}).Result()
	if err != nil {
		return fmt.Errorf("failed to add to sorted set: %w", err)
	}

	_, err = r.rdb.ZRemRangeByRank(context.TODO(), key, 0, -4).Result()
	if err != nil {
		return fmt.Errorf("failed to trim sorted set: %w", err)
	}

	return nil
}

func (r *redisCache) GetList(key string) ([]string, []int64, error) {
	valsWithTimestamps, err := r.rdb.ZRangeWithScores(context.TODO(), key, -3, -1).Result()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get sorted set: %w", err)
	}

	var values []string
	var timestamps []int64

	for _, z := range valsWithTimestamps {
		values = append(values, z.Member.(string))
		timestamps = append(timestamps, int64(z.Score))
	}

	return values, timestamps, nil
}
