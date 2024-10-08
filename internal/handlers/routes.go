package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/romacardozx/DEUNA-Challenge/internal/handlers/middleware"
	v1 "github.com/romacardozx/DEUNA-Challenge/internal/handlers/v1"
)

func SetupRoutes(router *gin.Engine, paymentHandler *v1.PaymentHandler, refundHandler *v1.RefundHandler) {
	router.GET("/health", healthCheck)

	api := router.Group("/api/v1")
	{
		auth := api.Group("/")
		auth.Use(middleware.AuthMiddleware())
		// Payment routes
		auth.POST("/payment", paymentHandler.ProcessPayment)
		auth.GET("/payment/details/:id", paymentHandler.GetPaymentDetails)
		auth.GET("/merchant/:merchantId/payments", paymentHandler.ListMerchantPayments)

		// Refund routes
		auth.POST("/refund", refundHandler.ProcessRefund)
		auth.GET("/refund/details/:refundId", refundHandler.GetRefundDetails)
	}

}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
