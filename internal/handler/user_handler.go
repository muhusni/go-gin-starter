package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	ID              uint64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name            string     `json:"name" gorm:"column:name;type:varchar(255)"`
	Username        string     `json:"username" gorm:"column:username;type:varchar(255)"`
	Email           string     `json:"email" gorm:"column:email;type:varchar(255)"`
	EmailVerifiedAt *time.Time `json:"email_verified_at" gorm:"column:email_verified_at"`
	Password        string     `json:"password" gorm:"column:password;type:varchar(255)"`
	RememberToken   string     `json:"remember_token" gorm:"column:remember_token;type:varchar(100)"`
	IsAdmin         bool       `json:"is_admin" gorm:"column:is_admin;type:tinyint(1)"`
	CreatedAt       time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"column:updated_at"`
}
type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []User

	err := h.DB.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
	}

	var user User
	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
	}

	var user User
	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
	}
}
