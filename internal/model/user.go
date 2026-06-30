package model

import (
	"time"
)

type User struct {
	ID              uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name            string     `json:"name" gorm:"column:name;type:varchar(255)"`
	Email           string     `json:"email" gorm:"column:email;type:varchar(255)"`
	EmailVerifiedAt *time.Time `json:"email_verified_at" gorm:"column:email_verified_at"`
	Password        string     `json:"password" gorm:"column:password;type:varchar(255)"`
	RememberToken   string     `json:"remember_token" gorm:"column:remember_token;type:varchar(100)"`
	IsAdmin         bool       `json:"is_admin" gorm:"column:is_admin;type:tinyint(1)"`
	CreatedAt       time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"column:updated_at"`
}
