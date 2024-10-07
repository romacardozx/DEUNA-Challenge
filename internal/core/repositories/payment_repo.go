package repositories

import "github.com/romacardozx/DEUNA-Challenge/internal/core/models"

func CreatePayment( /* parámetros necesarios */ ) (*models.Payment, error) {
	// Aquí iría la lógica para insertar en la base de datos
	// ...
	return &models.Payment{ /* campos */ }, nil
}

func GetPayment(id string) (*models.Payment, error) {
	// Aquí iría la lógica para obtener de la base de datos
	// ...
	return &models.Payment{ /* campos */ }, nil
}

func GetMerchantPayments(merchantID string) ([]*models.Payment, error) {
	// Aquí iría la lógica para obtener los pagos de un comerciante
	// ...
	return []*models.Payment{ /* lista de pagos */ }, nil
}
