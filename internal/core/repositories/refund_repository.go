package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
	"github.com/romacardozx/DEUNA-Challenge/internal/database"
)

type RefundRepository interface {
	Create(refund *models.Refund) error
	GetByID(refundID string) (*models.Refund, error)
}

var ErrRefundAlreadyExists = errors.New("refund already exists for this payment")

type refundRepository struct{}

func NewRefundRepository() RefundRepository {
	return &refundRepository{}
}

func (r *refundRepository) Create(refund *models.Refund) error {
	db := database.GetDB()
	query := `
		INSERT INTO refunds (id, payment_id, amount, currency, reason, status)
		SELECT $1, $2, $3, $4, $5, $6
		WHERE NOT EXISTS (
			SELECT id FROM refunds WHERE payment_id = $7
		)
		RETURNING id`

	var returnedID string
	err := db.QueryRow(query, refund.ID, refund.PaymentID, refund.Amount, refund.Currency, refund.Reason, refund.Status, refund.PaymentID).Scan(&returnedID)

	if err != nil {
		if err == sql.ErrNoRows {
			return ErrRefundAlreadyExists
		}
		return fmt.Errorf("error creating refund: %w", err)
	}

	if returnedID == "" {
		return ErrRefundAlreadyExists
	}

	return nil
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
