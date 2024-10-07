package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/romacardozx/DEUNA-Challenge/config"
	"github.com/romacardozx/DEUNA-Challenge/internal/database"
	"github.com/romacardozx/DEUNA-Challenge/internal/handlers"
)

type App struct {
	router *gin.Engine
	config *config.Config
	db     *database.Database
}

func NewApp() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	log.Printf("Config loaded. Migrations Path: %s", cfg.MigrationsPath)

	if err := database.RunMigrations(cfg.DatabaseURL, cfg.MigrationsPath); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	db, err := database.Init(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	router := gin.Default()

	app := &App{
		router: router,
		config: cfg,
		db:     db,
	}

	handlers.SetupRoutes(router, app.db)

	return app, nil
}

func (a *App) Run() error {
	return a.router.Run(a.config.ServerAddress)
}
