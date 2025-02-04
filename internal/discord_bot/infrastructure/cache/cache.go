package cache

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Add(key, value string) error
	GetArr(key string) ([]string, []int64, error)
}
