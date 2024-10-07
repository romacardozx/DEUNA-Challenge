package models

import "time"

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

type Payment struct {
	ID          string        `json:"id"`
	CustomerID  string        `json:"customer_id"`
	MerchantID  string        `json:"merchant_id"`
	Amount      float64       `json:"amount"`
	Currency    string        `json:"currency"`
	Description string        `json:"description"`
	Status      PaymentStatus `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
