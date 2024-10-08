package models

import (
	"time"

	"github.com/google/uuid"
)

type Refund struct {
	ID            uuid.UUID `json:"id"`
	PaymentID     string    `json:"payment_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Reason        string    `json:"reason"`
	Status        string    `json:"status"`
	TransactionID string    `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RefundPayload struct {
	PaymentID string
	Reason    string `json:"reason"`
}
