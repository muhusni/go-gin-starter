package app

import (
	"github.com/muhusni/go-gin-starter/internal/config"
	"github.com/muhusni/go-gin-starter/internal/database"
	"github.com/muhusni/go-gin-starter/internal/router"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	db, err := database.Connect(cfg)
	if err != nil {
		return err
	}
	r := router.New(db)
	return r.Run(":" + cfg.Port)
}
