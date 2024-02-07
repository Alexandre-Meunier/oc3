package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func newRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     viper.GetString("Redis.Address"),
		Password: viper.GetString("Redis.Password"),
		DB:       viper.GetInt("Redis.Database"),
	})
}
