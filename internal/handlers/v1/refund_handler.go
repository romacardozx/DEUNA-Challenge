package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romacardozx/DEUNA-Challenge/internal/database"
)

func CreateRefund(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Refund created successfully",
			"refundId":  "mock-refund-id-123",
			"paymentId": "mock-payment-id-123",
			"amount":    50.00,
			"currency":  "USD",
			"status":    "processed",
		})
	}
}
