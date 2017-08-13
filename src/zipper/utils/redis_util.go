package utils

import (
	"gopkg.in/redis.v3"
	"time"
	"errors"
	"fmt"
	"zipper/config"
	"strconv"
)

var (
	redisClient *RedisClient
)

type RedisClient struct {
	client *redis.Client
}

func (this *RedisClient) SaveKey(key string, value []byte, duration time.Duration) (err error) {
	if err = this.client.Set(key, value, duration).Err(); err != nil {
		err = errors.New(fmt.Sprintf("Error when saving to redis. key: %s, error: %s",
			key, err.Error()))
	}
	return
}

func (this *RedisClient) GetValue(key string) (value []byte, err error) {
	if value, err = this.client.Get(key).Bytes(); err != nil {
		err = errors.New(fmt.Sprintf("Error when getting from redis. key: %s, error: %s",
			key, err.Error()))
	}
	return
}

func GetRedisClient() *RedisClient {
	if redisClient == nil {
		conf := config.Redis
		db, _ := strconv.ParseInt(conf.DB, 10, 64)

		rc := redis.NewClient(&redis.Options{
			Addr:     conf.Address,
			Password: conf.Password,
			DB:       db,
			MaxRetries: 1,
			IdleTimeout: 1 * time.Minute,
		})
		redisClient = &RedisClient{client:rc}
	}
	return redisClient
}