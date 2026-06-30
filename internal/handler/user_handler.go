package handler

import (
	"net/http"
	"strconv"
	"github.com/muhusni/go-gin-starter/internal/model"
	"github.com/muhusni/go-gin-starter/internal/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []model.User

	err := h.DB.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return 
	}

	responses := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, dto.UserResponse{
			ID: user.ID,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			IsAdmin: user.IsAdmin,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": responses,
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

	var user model.User
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

	var user model.User
	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

}

func ToUserResponse(user model.User) dto.UserResponse {
	return dto.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Username: user.Username,
		Email: user.Email,
		IsAdmin: user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
