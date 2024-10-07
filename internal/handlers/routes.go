package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/romacardozx/DEUNA-Challenge/internal/database"
	v1 "github.com/romacardozx/DEUNA-Challenge/internal/handlers/v1"
)

func SetupRoutes(router *gin.Engine, db *database.Database) {
	router.GET("/health", healthCheck)

	apiV1 := router.Group("/api/v1")
	{
		payments := apiV1.Group("/payments")
		{
			payments.POST("", v1.CreatePayment(db))
			payments.GET("/:id", v1.GetPayment(db))
		}

		refunds := apiV1.Group("/refunds")
		{
			refunds.POST("", v1.CreateRefund(db))
		}

		merchants := apiV1.Group("/merchants")
		{
			merchants.GET("/:id/payments", v1.GetMerchantPayments(db))
		}
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
