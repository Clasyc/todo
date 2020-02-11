package model

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

var svc Service

func InitService(db *gorm.DB) {
	svc = Service{
		DB:db,
	}
}