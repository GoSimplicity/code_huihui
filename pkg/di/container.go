package di

import (
	"github.com/GoSimplicity/code_huihui/internal/api"
	"github.com/GoSimplicity/code_huihui/internal/repository"
	"github.com/GoSimplicity/code_huihui/internal/repository/dao"
	"github.com/GoSimplicity/code_huihui/internal/service"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitContainer() *do.Injector {
	injector := do.New()
	do.Provide(injector, InitViper)
	do.Provide(injector, InitLogger)
	do.Provide(injector, InitDB)
	do.Provide(injector, InitRedis)
	do.Provide(injector, dao.NewUserDAO)
	do.Provide(injector, repository.NewUserRepository)
	do.Provide(injector, service.NewUserService)
	do.Provide(injector, api.NewUserHandler)

	return injector
}

func MustInitContainer() *do.Injector {
	return InitContainer()
}

func GetLogger(i *do.Injector) *zap.Logger {
	return do.MustInvoke[*zap.Logger](i)
}

func GetDB(i *do.Injector) *gorm.DB {
	return do.MustInvoke[*gorm.DB](i)
}

func GetRedis(i *do.Injector) redis.Cmdable {
	return do.MustInvoke[redis.Cmdable](i)
}

func GetUserService(i *do.Injector) service.UserService {
	return do.MustInvoke[service.UserService](i)
}

func GetUserHandler(i *do.Injector) *api.UserHandler {
	return do.MustInvoke[*api.UserHandler](i)
}

func GetViper(i *do.Injector) *viper.Viper {
	return do.MustInvoke[*viper.Viper](i)
}

func ShutdownContainer(i *do.Injector) error {
	return i.Shutdown()
}
