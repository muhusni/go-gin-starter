package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muhusni/go-gin-starter/internal/config"
	"github.com/muhusni/go-gin-starter/internal/handler"
	"github.com/muhusni/go-gin-starter/internal/middleware"
	"github.com/muhusni/go-gin-starter/internal/repository"
	"github.com/muhusni/go-gin-starter/internal/security"
	"github.com/muhusni/go-gin-starter/internal/service"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()
	// repositories
	userRepository := repository.NewUserRepository(db)

	// Services
	jwtService := security.NewJWTService(cfg.JWT.Secret, cfg.JWT.Issuer)
	authService := service.NewAuthService(jwtService, userRepository)
	userService := service.NewUserService(userRepository)

	// handlers
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

	// middleware
	mw := middleware.NewMiddleware(jwtService)

	// routes
	r.GET("/favicon.ico", handler.FaviconHandler)
	r.GET("/ping", handler.PingHandler)
	r.GET("/health", handler.HealthHandler)

	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	protected := api.Group("/")
	protected.Use(mw.Auth())
	{
		protected.GET("/me", authHandler.Me)
		user := protected.Group("/users")
		{
			user.GET("", userHandler.GetUsers)
			user.POST("", userHandler.CreateUser)
			user.GET("/:id", userHandler.GetUser)
			user.PUT("/:id", userHandler.UpdateUser)
			user.DELETE("/:id", userHandler.DeleteUser)
		}
	}
	return r
}
