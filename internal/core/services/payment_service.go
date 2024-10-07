package services

import (
	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/repositories"
)

func CreatePayment( /* parámetros necesarios */ ) (*models.Payment, error) {
	// Lógica de negocio aquí
	// ...

	// Llamada al repositorio
	payment, err := repositories.CreatePayment( /* parámetros */ )
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func GetPayment(id string) (*models.Payment, error) {
	return repositories.GetPayment(id)
}

func GetMerchantPayments(merchantID string) ([]*models.Payment, error) {
	return repositories.GetMerchantPayments(merchantID)
}
