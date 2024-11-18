// 初始化数据库
package config

import (
	"SUPPORT/global"
	"SUPPORT/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() { // 初始化数据库
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxidleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxopenConns)
	sqlDB.SetConnMaxLifetime(time.Hour) // 1 hour

	if err != nil {
		log.Fatalf("failed to set database connection: %v", err)
	}

	global.Db = db //为其他包提供数据库连接
	db.AutoMigrate(&models.Load{})
}
