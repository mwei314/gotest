package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisClient redis客户端
var RedisClient *redis.Pool

// Init 初始化redis客户端
func Init() {
	// 建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:     5,
		MaxActive:   10,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", ":6379",
				redis.DialPassword(""),
				redis.DialDatabase(0),
				redis.DialConnectTimeout(20*time.Second),
				redis.DialReadTimeout(10*time.Second),
				redis.DialWriteTimeout(10*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}

// Do 执行redis命令
func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	if RedisClient == nil {
		panic("redis error")
	}
	conn := RedisClient.Get()
	defer conn.Close()
	return conn.Do(commandName, args...)
}
