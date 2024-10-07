package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/services"
)

func CreateRefund(c *gin.Context) {
	// Aquí iría la lógica para extraer datos de la solicitud
	// Por ejemplo:
	// var refundRequest RefundRequest
	// if err := c.ShouldBindJSON(&refundRequest); err != nil {
	//     c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//     return
	// }

	// Llamada al servicio
	refund, err := services.CreateRefund("", 0, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, refund)
}
