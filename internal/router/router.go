package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muhusni/go-gin-starter/internal/handler"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	userHandler := handler.NewUserHandler(db)

	r.GET("/favicon.ico", handler.FaviconHandler)
	r.GET("/ping", handler.PingHandler)
	r.GET("/health", handler.HealthHandler)

	apiUser := r.Group("/users")
	apiUser.GET("/", userHandler.GetUsers)
	apiUser.GET("/:id", userHandler.GetUser)
	return r
}
