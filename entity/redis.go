package entity

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func SetKey(key string, value []byte) error {
	ctx := context.Background()
	rdb := InitRedis()
	err := rdb.Set(ctx, key, value, 600).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetKey(key string) ([]byte, error) {
	ctx := context.Background()
	rdb := InitRedis()
	val, err := rdb.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return val, nil
}
