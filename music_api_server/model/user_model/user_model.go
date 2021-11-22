package user_model

import (
	"gorm.io/gorm"
)

func init(){}

type (
	UserInfo struct {
		gorm.Model
		Email string `gorm:"email;SIZE:32;NOT NULL"`
		Password string `gorm:"passoword;size:64;NOT NULL"`
		FirstName string `gorm:"name;size:16;default:user_service;NOT NULL" `
		LastName string  `gorm:"name;size:16;default:user_service;NOT NULL" `
		Icon string `gorm:"image;size:32;NOT NULL"`
	}
)
