package dto

import "time"

type UserResponse struct {
	ID           	uint64		`json:"name"`
	Name            string		`json:"name"`
	Username        string		`json:"username"`
	Email           string		`json:"email"`
	Password        string		`json:"password"`
	IsAdmin         bool  		`json:"is_admin"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}