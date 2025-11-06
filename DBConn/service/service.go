package service

import (
	"strconv"
	"yourapp/db"
	"yourapp/logger"
	"yourapp/model"

	"go.uber.org/zap"
)

func GetUser(id int) (*model.User, error) {
	var user model.User
	err := db.DB.First(&user, id).Error
	return &user, err
}

func InsertUser(user model.User) {

	result := db.DB.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	logger.Log.Info("Inserted user ID:", zap.String("userid", strconv.FormatUint(uint64(user.ID), 10)))

}
