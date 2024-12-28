package service

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"web/model"
	"web/utils/config"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.Dbname,
	)), &gorm.Config{})

	if err != nil {
		println("数据库连接失败：", err)
	}

	// todo:数据库迁移，创建需要的表
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserCard{})

	sqlDB, err := db.DB()
	if err != nil {
		println("获取通用数据库对象 sql.DB 出错：", err)
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
