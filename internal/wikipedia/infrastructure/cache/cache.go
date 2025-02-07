package cache

import "time"

type Cache interface {
	SetS(key, value string) error
	GetS(key string) (string, error)
	AddS(key, value string) error
	GetList(key string) ([]string, []int64, error)
	GetZRangeByDate(date time.Time, key string) ([]string, []int64, error)
}
