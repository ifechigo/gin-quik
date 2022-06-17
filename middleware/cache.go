package middleware

import (
	"fmt"
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()


func PostCache(key string, value float64)  {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	
}

func GetCache(key string)  {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,})
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("Wallet with ID:%s does not exist", key)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("Wallet ID:%s \nBalance:%s", key, val)
	}
}