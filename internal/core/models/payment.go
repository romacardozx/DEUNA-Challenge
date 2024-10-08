package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	PaymentStatusPending   string = "pending"
	PaymentStatusCompleted string = "completed"
	PaymentStatusFailed    string = "failed"
	PaymentStatusRefunded  string = "refunded"
)

type Payment struct {
	ID            uuid.UUID `json:"id"`
	CustomerID    string    `json:"customer_id"`
	MerchantID    string    `json:"merchant_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	TransactionID string    `json:"transaction_id"`
}

type PaymentPayload struct {
	CustomerID  string  `json:"customer_id"`
	MerchantID  string  `json:"merchant_id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
}
