package main

import (
	"fmt"
	"github.com/go-redis/redis"
	redisCli "github.com/kitakitabauer/go-redis-sample/redis"

)

func main() {
	fmt.Println("test")
	cli := redisCli.NewRedis()

	err := cli.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := cli.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := cli.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
