package repositories

import (
	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
	"github.com/romacardozx/DEUNA-Challenge/internal/database"
)

type RefundRepository interface {
	Create(refund *models.Refund) error
	GetByID(refundID string) (*models.Refund, error)
	ListByPayment(paymentID string) ([]*models.Refund, error)
}

type refundRepository struct{}

func NewRefundRepository() RefundRepository {
	return &refundRepository{}
}

func (r *refundRepository) Create(refund *models.Refund) error {
	db := database.GetDB()
	query := `INSERT INTO refunds (id, payment_id, amount, currency, reason, status) 
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, refund.ID, refund.PaymentID, refund.Amount, refund.Currency, refund.Reason, refund.Status)
	return err
}

func (r *refundRepository) GetByID(refundID string) (*models.Refund, error) {
	db := database.GetDB()
	query := `SELECT id, payment_id, amount, currency, reason, status FROM refunds WHERE id = $1`
	refund := &models.Refund{}
	err := db.QueryRow(query, refundID).Scan(&refund.ID, &refund.PaymentID, &refund.Amount, &refund.Currency, &refund.Reason, &refund.Status)
	if err != nil {
		return nil, err
	}
	return refund, nil
}

func (r *refundRepository) ListByPayment(paymentID string) ([]*models.Refund, error) {
	db := database.GetDB()
	query := `SELECT id, payment_id, amount, currency, reason, status FROM refunds WHERE payment_id = $1`
	rows, err := db.Query(query, paymentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var refunds []*models.Refund
	for rows.Next() {
		refund := &models.Refund{}
		err := rows.Scan(&refund.ID, &refund.PaymentID, &refund.Amount, &refund.Currency, &refund.Reason, &refund.Status)
		if err != nil {
			return nil, err
		}
		refunds = append(refunds, refund)
	}
	return refunds, nil
}
