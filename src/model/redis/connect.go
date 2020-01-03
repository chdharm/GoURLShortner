package redisconnect

import (
	"github.com/go-redis/redis"
)

func Connect(){
	/*------- Redis Config ----------*/
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	redisPong, redisErr := redisClient.Ping().Result()
	if redisErr != nil {
		panic(redisErr)
	}
	if redisPong == "PONG" {
		fmt.Println("Redis client connected")
	}
	/*------------------------------------*/
}
