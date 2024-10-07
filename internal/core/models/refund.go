package models

import "time"

type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusCompleted RefundStatus = "completed"
	RefundStatusFailed    RefundStatus = "failed"
)

type Refund struct {
	ID          string       `json:"id"`
	PaymentID   string       `json:"payment_id"`
	Amount      float64      `json:"amount"`
	Currency    string       `json:"currency"`
	Reason      string       `json:"reason"`
	Status      RefundStatus `json:"status"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
