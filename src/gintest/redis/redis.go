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
				redis.DialConnectTimeout(5*time.Second),
				redis.DialReadTimeout(2*time.Second),
				redis.DialWriteTimeout(2*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}

// Do 执行redis命令
func Do(commandName string, args ...interface{}) (interface{}, error) {
	if RedisClient == nil {
		panic("redis error")
	}
	conn := RedisClient.Get()
	defer conn.Close()
	return conn.Do(commandName, args...)
}

// MultiDo 管道执行命令
func MultiDo(para [][]interface{}) ([]interface{}, error) {
	conn := RedisClient.Get()
	defer conn.Close()
	for _, v := range para {
		conn.Send(v[0].(string), v[1:]...)
	}
	return redis.Values(conn.Do("Exec"))
}
