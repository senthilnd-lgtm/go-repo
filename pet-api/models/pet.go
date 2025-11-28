package models

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}
