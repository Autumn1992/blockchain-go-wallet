package db

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type REDISDB int

// 定义多个 redis db实例 功能分开
const (
	COMMON   REDISDB = 10 // 通用
	BANKDATA REDISDB = 15 // 银行数据
)

var redisMap map[REDISDB]*redis.Client

func initRedis() {
	redisMap = make(map[REDISDB]*redis.Client)

	for _, redisConf := range GetConfig().Redis {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     redisConf.Host,
			Password: redisConf.Pwd,
			DB:       redisConf.Db,
			//IdleTimeout: 300, // 默认Idle超时时间
			PoolSize: 100, // 连接池
		})
		_, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			panic(err)
		}
		redisMap[REDISDB(redisConf.Db)] = redisClient
	}
}

func GetRedis() *redis.Client {
	return redisMap[COMMON]
}

func GetAgencyRedis() *redis.Client {
	return GetRedis()
}
func GetMaxWinnerRedis() *redis.Client {
	return GetRedis()
}

func GetGameTagRedis() *redis.Client {
	return GetRedis()
}

func GetBankRedis() *redis.Client {
	return redisMap[BANKDATA]
}
