package main

import (
	"yourapp/logger"
	"yourapp/service"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Logger initialized")
	service.DoSomething()
}
