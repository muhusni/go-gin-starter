package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muhusni/go-gin-starter/internal/handler"
	"github.com/muhusni/go-gin-starter/internal/service"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler()
	r.GET("/favicon.ico", handler.FaviconHandler)
	r.GET("/ping", handler.PingHandler)
	r.GET("/health", handler.HealthHandler)

	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	user := api.Group("/users")
	{
		user.GET("", userHandler.GetUsers)
		user.POST("", userHandler.CreateUser)
		user.GET("/:id", userHandler.GetUser)
		user.PUT("/:id", userHandler.UpdateUser)
		user.DELETE("/:id", userHandler.DeleteUser)
	}
	return r
}
