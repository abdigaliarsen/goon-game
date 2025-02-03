package utils

import (
	"log"
	"os"
	"strconv"
	"time"
)

func MustGetEnv[T any](key string) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Missing required environment variable: %s", key)
	}

	var zero T
	switch any(zero).(type) {
	case string:
		return any(val).(T)

	case int:
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("Environment variable %s must be an int (got %q): %v", key, val, err)
		}
		return any(i).(T)

	case bool:
		b, err := strconv.ParseBool(val)
		if err != nil {
			log.Fatalf("Environment variable %s must be a bool (got %q): %v", key, val, err)
		}
		return any(b).(T)

	case time.Duration:
		d, err := time.ParseDuration(val)
		if err != nil {
			log.Fatalf("Environment variable %s must be a valid duration (got %q): %v", key, val, err)
		}
		return any(d).(T)

	default:
		log.Fatalf("Unsupported type for environment variable parsing: %T", zero)
	}

	return zero
}
