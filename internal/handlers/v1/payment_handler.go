package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romacardozx/DEUNA-Challenge/internal/database"
)

func CreatePayment(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Payment created successfully",
			"paymentId": "mock-payment-id-123",
		})
	}
}

func GetPayment(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		paymentId := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"paymentId": paymentId,
			"amount":    100.00,
			"currency":  "USD",
			"status":    "completed",
		})
	}
}

func GetMerchantPayments(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		merchantId := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"merchantId": merchantId,
			"payments": []gin.H{
				{
					"paymentId": "mock-payment-id-1",
					"amount":    100.00,
					"currency":  "USD",
					"status":    "completed",
				},
				{
					"paymentId": "mock-payment-id-2",
					"amount":    150.00,
					"currency":  "USD",
					"status":    "pending",
				},
			},
		})
	}
}
