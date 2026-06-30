package service

import "gorm.io/gorm"

type AuthService struct {
	DB *gorm.DB
}
