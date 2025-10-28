package dbopr

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
	fmt.Println("Called")

	// running as docker container
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/my_database?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println("DB starting on port 3306...")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to MySQL and auto-migrated!")

	return db

}
