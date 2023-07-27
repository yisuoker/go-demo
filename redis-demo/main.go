package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yisuoker/go-demo/conf"
)

var rdb *redis.Client

func getResult() {
	// Context 上下文
	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var key string
	key = "hello"

	// 执行命令获取结果
	val, err := rdb.Get(ctx, key).Result()

	// redis.Nil 表示 Key 不存在
	if err != nil {
		if errors.Is(err, redis.Nil) {
			fmt.Println(key + "不存在")
			return
		}
		// 其他错误
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func get() {
	// Context 上下文
	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var key string
	key = "hello"

	// 先获取执行命令对象，再获取结果（值和错误）
	cmder := rdb.Get(ctx, key)
	fmt.Println(cmder.Val())
	fmt.Println(cmder.Err())
}

func set() {
	// Context 上下文
	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var key string
	key = "hello"

	if err := rdb.Set(ctx, key, "welcome to redis world!", time.Second*5).Err(); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	if err := conf.Init(""); err != nil {
		fmt.Printf("Failed to load the configuration file, err: %v\n", err)
		return
	}
	// fmt.Printf("%#v\n", conf.Cfg)

	cfg := conf.Cfg.RedisCfg
	// fmt.Printf("%#v\n", cfg)

	Addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Username: "",
		Password: "",
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	set()
	// get()
	getResult()
}
