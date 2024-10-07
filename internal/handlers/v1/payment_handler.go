package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/services"
)

func CreatePayment(c *gin.Context) {
	// Aquí iría la lógica para extraer datos de la solicitud
	// Por ejemplo:
	// var paymentRequest PaymentRequest
	// if err := c.ShouldBindJSON(&paymentRequest); err != nil {
	//     c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//     return
	// }

	// Llamada al servicio
	payment, err := services.CreatePayment( /* pasar los datos necesarios */ )
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func GetPayment(c *gin.Context) {
	id := c.Param("id")

	payment, err := services.GetPayment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func GetMerchantPayments(c *gin.Context) {
	merchantID := c.Param("id")

	payments, err := services.GetMerchantPayments(merchantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}
