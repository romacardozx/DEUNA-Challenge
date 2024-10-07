package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/romacardozx/online-payment-platform/internal/handlers"
)

type App struct {
	router *gin.Engine
}

func NewApp() (*App, error) {
	router := gin.Default()

	handlers.SetupRoutes(router)

	return &App{
		router: router,
	}, nil
}

func (a *App) Run() error {
	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	return a.router.Run(port)
}
