package service

import (
	"yourapp/logger"

	"go.uber.org/zap"
)

func DoSomething() {
	logger.Log.Info("Doing something", zap.String("module", "service"))
}
