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

	r.GET("/favicon.ico", handler.FaviconHandler)
	r.GET("/ping", handler.PingHandler)
	r.GET("/health", handler.HealthHandler)

	api := r.Group("/api/v1")

	apiUser := api.Group("/users")
	{
		apiUser.GET("", userHandler.GetUsers)
		apiUser.POST("", userHandler.CreateUser)
		apiUser.GET("/:id", userHandler.GetUser)
		apiUser.PUT("/:id", userHandler.UpdateUser)
		apiUser.DELETE("/:id", userHandler.DeleteUser)
	}
	return r
}
