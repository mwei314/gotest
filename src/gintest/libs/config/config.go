package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

type (
	// Config 配置文件数据
	Config struct {
		Test   string
		Redis  Redis
		Mysql  string
		Logger Logger
	}

	// Redis redis客户端配置
	Redis struct {
		Addr     string
		Password string
		DB       int `toml:"db"`
	}

	// Logger 日志相关配置
	Logger struct {
		Path string
	}
)

var (
	cfg  *Config
	once sync.Once
)

// Init 初始化配置
func Init() *Config {
	once.Do(func() {
		confPath := getConfFile()
		filePath, err := filepath.Abs(confPath)
		if err != nil {
			panic(err)
		}
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}

// getConfFile 根据环境变量设置构造配置文件路径
func getConfFile() string {
	environment, ok := os.LookupEnv("GOLANG_ENVIRONMENT")
	if !ok {
		// 没有设置环境变量 使用dev
		environment = "dev"
	}
	return "conf/" + environment + ".toml"
}
