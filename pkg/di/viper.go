package di

import (
	"fmt"

	"github.com/samber/do"
	"github.com/spf13/viper"
)

func InitViper(_ *do.Injector) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath(".")
	v.SetDefault("server.port", "8080")
	v.SetDefault("server.mode", "debug")
	v.SetDefault("log.level", "info")
	v.SetDefault("log.dir", "./logs")
	v.SetDefault("mysql.addr", "root:password@tcp(localhost:3306)/code_huihui?charset=utf8mb4&parseTime=True&loc=Local")
	v.SetDefault("redis.addr", "localhost:6379")
	v.SetDefault("redis.password", "")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Warning: Config file not found, using defaults: %v\n", err)
	}

	return v, nil
}
