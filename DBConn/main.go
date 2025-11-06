package main

import (
	"fmt"
	"yourapp/db"
	"yourapp/logger"
	"yourapp/model"
	"yourapp/service"
)

func main() {

	// Initialize logger

	logger.InitLogger()
	logger.Log.Info("DB initialized")

	// Initialize DB

	db.InitMySQL()
	logger.Log.Info("Logger initialized")

	user1 := model.User{Name: "Senthilkumar", Email: "sen7657@gmail.com"}
	service.InsertUser(user1)
	fmt.Println(service.GetUser(1))

}
