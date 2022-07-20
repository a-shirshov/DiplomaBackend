package utils

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"errors"
)

func InitRedisDB() (*redis.Client, error) {
	addr := viper.GetString("redis.addr")
	dbId := viper.GetInt("redis.db_id")
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: "",
		DB: dbId,
	})
	
	if client == nil {
		return nil, errors.New("redis db init failed")
	}

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}