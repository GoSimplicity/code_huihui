package di

import (
	"fmt"
	"log"

	"github.com/samber/do"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(i *do.Injector) (*gorm.DB, error) {
	v := do.MustInvoke[*viper.Viper](i)
	addr := v.GetString("mysql.addr")
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		log.Printf("数据库连接失败: %v", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("获取sql.DB失败: %v", err)
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	if err := InitTables(db); err != nil {
		log.Printf("初始化数据库表失败: %v", err)
	}
	return db, nil
}

func InitTables(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("数据库连接为空，跳过表初始化")
	}
	// TODO: 初始化表
	return nil
}
