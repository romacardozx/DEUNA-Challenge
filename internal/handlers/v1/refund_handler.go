package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/services"
)

type RefundHandler struct {
	refundService services.RefundService
}

func NewRefundHandler(refundService services.RefundService) *RefundHandler {
	return &RefundHandler{
		refundService: refundService,
	}
}

func (h *RefundHandler) ProcessRefund(c *gin.Context) {
	var refund models.Refund
	if err := c.ShouldBindJSON(&refund); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	processedRefund, err := h.refundService.ProcessRefund(&refund)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process refund"})
		return
	}

	c.JSON(http.StatusOK, processedRefund)
}

func (h *RefundHandler) GetRefundDetails(c *gin.Context) {
	refundID := c.Param("id")

	refund, err := h.refundService.GetRefundDetails(refundID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Refund not found"})
		return
	}

	c.JSON(http.StatusOK, refund)
}

func (h *RefundHandler) ListPaymentRefunds(c *gin.Context) {
	paymentID := c.Param("paymentId")

	refunds, err := h.refundService.ListPaymentRefunds(paymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve refunds"})
		return
	}

	c.JSON(http.StatusOK, refunds)
}
