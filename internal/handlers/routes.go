package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/romacardozx/DEUNA-Challenge/internal/database"
	v1 "github.com/romacardozx/DEUNA-Challenge/internal/handlers/v1"
)

func SetupRoutes(router *gin.Engine, db *database.Database) {
	router.GET("/health", healthCheck)

	api := router.Group("/api/v1")
	{
		api.POST("/payments", v1.CreatePayment)
		api.GET("/payments/:id", v1.GetPayment)
		api.POST("/refunds", v1.CreateRefund)
		api.GET("/merchants/:id/payments", v1.GetMerchantPayments)
	}

}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
