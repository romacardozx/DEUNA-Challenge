package app

import (
	"github.com/gin-gonic/gin"
	"github.com/romacardozx/DEUNA-Challenge/config"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/repositories"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/services"
	"github.com/romacardozx/DEUNA-Challenge/internal/database"
	"github.com/romacardozx/DEUNA-Challenge/internal/handlers"
	v1 "github.com/romacardozx/DEUNA-Challenge/internal/handlers/v1"
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

	db, err := database.Init(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	router := gin.Default()

	// Initialize repositories
	paymentRepo := repositories.NewPaymentRepository()
	refundRepo := repositories.NewRefundRepository()

	// Initialize services
	bankSimulator := services.NewBankSimulatorService()
	paymentService := services.NewPaymentService(paymentRepo, bankSimulator)
	refundService := services.NewRefundService(refundRepo, paymentRepo, bankSimulator)

	// Initialize handlers
	paymentHandler := v1.NewPaymentHandler(paymentService)
	refundHandler := v1.NewRefundHandler(refundService)

	app := &App{
		router: router,
		config: cfg,
		db:     db,
	}

	handlers.SetupRoutes(router, paymentHandler, refundHandler)

	return app, nil
}

func (a *App) Run() error {
	return a.router.Run(a.config.ServerAddress)
}
