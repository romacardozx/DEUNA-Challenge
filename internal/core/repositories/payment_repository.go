package repositories

import (
	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
	"github.com/romacardozx/DEUNA-Challenge/internal/database"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetByID(paymentID string) (*models.Payment, error)
	ListByMerchant(merchantID string, limit, offset int) ([]models.Payment, error)
}

type paymentRepository struct{}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepository{}
}

func (r *paymentRepository) Create(payment *models.Payment) error {
	db := database.GetDB()
	query := `INSERT INTO payments (id, amount, currency, customer_id, merchant_id, status, transaction_id, description, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := db.Exec(query, payment.ID, payment.Amount, payment.Currency, payment.CustomerID, payment.MerchantID, payment.Status, payment.TransactionID, payment.Description, payment.CreatedAt)
	return err
}

func (r *paymentRepository) GetByID(paymentID string) (*models.Payment, error) {
	db := database.GetDB()
	query := `SELECT id, amount, currency, customer_id, merchant_id, status, transaction_id, created_at, updated_at, description FROM payments WHERE id = $1`
	payment := &models.Payment{}
	err := db.QueryRow(query, paymentID).Scan(&payment.ID, &payment.Amount, &payment.Currency, &payment.CustomerID, &payment.MerchantID, &payment.Status, &payment.TransactionID, &payment.CreatedAt, &payment.UpdatedAt, &payment.Description)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *paymentRepository) ListByMerchant(merchantID string, limit, offset int) ([]models.Payment, error) {
	db := database.GetDB()
	payments := []models.Payment{}
	query := `SELECT id, amount, currency, customer_id, merchant_id, status, transaction_id, created_at, updated_at, description 
              FROM payments WHERE merchant_id = $1 LIMIT $2 OFFSET $3`
	rows, err := db.Query(query, merchantID, limit, offset)
	if err != nil {
		return payments, err
	}
	defer rows.Close()

	for rows.Next() {
		payment := models.Payment{}
		err := rows.Scan(&payment.ID, &payment.Amount, &payment.Currency, &payment.CustomerID, &payment.MerchantID, &payment.Status, &payment.TransactionID, &payment.CreatedAt, &payment.UpdatedAt, &payment.Description)
		if err != nil {
			return payments, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}
