package db

import (
	"time"
	"yourapp/logger"
	"yourapp/model"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL() {
	dsn := "root:my-secret-pw@tcp(192.168.29.167:3306)/my_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal("Failed to connect to DB", zap.Error(err))
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Log.Fatal("Failed to get sql.DB", zap.Error(err))
		return
	}
	db.AutoMigrate(&model.User{})

	// Connection pool settings
	sqlDB.SetMaxOpenConns(25)                 // max open connections
	sqlDB.SetMaxIdleConns(10)                 // max idle connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // max lifetime per connection

	DB = db
}
