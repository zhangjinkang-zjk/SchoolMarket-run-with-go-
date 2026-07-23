package database

import (
	"fmt"
	"time"

	"SchoolMarket-run-with-go-/config"
	"SchoolMarket-run-with-go-/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("数据库连接失败:%v\n", err))
	}

	if err := db.AutoMigrate(&model.User{}, &model.Goods{}, &model.Cart{}); err != nil {
		panic(fmt.Sprintf("自动迁移失败:%v\n", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("获取数据库数据失败%v\n", err))
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	fmt.Println("数据库连接成功，数据获取成功")

	return db
}
