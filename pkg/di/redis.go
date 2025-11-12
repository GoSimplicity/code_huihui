package di

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/samber/do"
	"github.com/spf13/viper"
)

func InitRedis(i *do.Injector) (redis.Cmdable, error) {
	v := do.MustInvoke[*viper.Viper](i)
	client := redis.NewClient(&redis.Options{
		Addr:     v.GetString("redis.addr"),
		Password: v.GetString("redis.password"),
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("redis 连接失败，请检查密码和主机地址:redis.addr:%v,redis.password:%v,err:%v", v.GetString("redis.addr"), v.GetString("redis.password"), err)
	}

	return client, nil
}
