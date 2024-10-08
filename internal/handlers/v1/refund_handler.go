package v1

import (
	"fmt"
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
	var refundPayload models.RefundPayload
	if err := c.ShouldBindJSON(&refundPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Refund Payload: ", refundPayload)

	processedRefund, err := h.refundService.ProcessRefund(c, &refundPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, processedRefund)
}

func (h *RefundHandler) GetRefundDetails(c *gin.Context) {
	refundID := c.Param("id")

	refund, err := h.refundService.GetRefundDetails(c, refundID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Refund not found"})
		return
	}

	c.JSON(http.StatusOK, refund)
}
