package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/services"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

func (controller *PaymentHandler) ProcessPayment(c *gin.Context) {
	var paymentPayload models.PaymentPayload
	if err := c.ShouldBindJSON(&paymentPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding: " + err.Error()})
		return
	}

	processedPayment, err := controller.paymentService.ProcessPayment(c, &paymentPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, processedPayment)
}

func (controller *PaymentHandler) GetPaymentDetails(c *gin.Context) {
	paymentID := c.Param("id")

	payment, err := controller.paymentService.GetPaymentDetails(c, paymentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func (controller *PaymentHandler) ListMerchantPayments(c *gin.Context) {
	merchantID := c.Param("merchantId")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	payments, err := controller.paymentService.ListMerchantPayments(c, merchantID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payments"})
		return
	}

	c.JSON(http.StatusOK, payments)
}
