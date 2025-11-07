package database

import (
	"log"
	"pet-api/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/plugin/prometheus"
)

var DB *gorm.DB

func InitMySQL() {
	dsn := "root:my-secret-pw@tcp(192.168.29.167:3306)/my_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB")
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB")
		return
	}

	// Add Prometheus plugin
	db.Use(prometheus.New(prometheus.Config{
		DBName:          "pets",
		RefreshInterval: 15,
		StartServer:     true,
		HTTPServerPort:  9000, // exposes DB metrics at :9000/metrics
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{VariableNames: []string{"Threads_running"}},
		},
	}))

	db.AutoMigrate(&models.Pet{})

	// Connection pool settings
	sqlDB.SetMaxOpenConns(25)                 // max open connections
	sqlDB.SetMaxIdleConns(10)                 // max idle connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // max lifetime per connection

	DB = db
}
