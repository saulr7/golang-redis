package db

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	rdb  *redis.Client
	once sync.Once
)

func ConnectRedis() {

	once.Do(func() {

		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

	})

}

func PoolRDB() *redis.Client {
	return rdb
}
