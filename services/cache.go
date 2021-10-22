package services

import (
	"sync"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var (
	onceCache sync.Once
	c         *cache.Cache
)

// NewCache - Создаем экземпляр КиШа
func NewCache(config *Config) *cache.Cache {
	onceCache.Do(func() {
		ring := redis.NewRing(&redis.RingOptions{
			Addrs: map[string]string{
				"localhost": ":6379",
			},
		})

		c = cache.New(&cache.Options{
			Redis:      ring,
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		})
	})
	return c
}
