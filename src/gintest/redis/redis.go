package redis

import (
	"gintest/utils"
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisClient redis客户端
var RedisClient *redis.Pool

// Init 初始化redis客户端
func Init(conf *utils.Config) {
	// 建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:     5,
		MaxActive:   10,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", conf.Redis.Addr,
				redis.DialPassword(conf.Redis.Password),
				redis.DialDatabase(conf.Redis.DB),
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
