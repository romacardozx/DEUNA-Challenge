package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditData struct {
	ID            uuid.UUID
	PaymentID     uuid.UUID
	Amount        float64
	MerchantID    string
	CustomerID    string
	Currency      string
	TransactionID string
	Status        string
	Message       string
	CreatedAt     time.Time
}

func (d *AuditData) CompletePaymentLog(payment Payment, response BankResponse) {
	d.ID = uuid.New()
	d.PaymentID = payment.ID
	d.Amount = payment.Amount
	d.MerchantID = payment.MerchantID
	d.CustomerID = payment.CustomerID
	d.Currency = payment.Currency
	d.TransactionID = response.TransactionID
	d.Status = payment.Status
	d.Message = response.Message
	d.CreatedAt = time.Now()

}

func (d *AuditData) CompleteRefundLog(refund Refund, response BankResponse, merchantID string, customerID string) {
	d.ID = uuid.New()
	d.PaymentID = refund.ID
	d.Amount = refund.Amount
	d.MerchantID = merchantID
	d.CustomerID = customerID
	d.Currency = refund.Currency
	d.TransactionID = response.TransactionID
	d.Status = refund.Status
	d.Message = response.Message
	d.CreatedAt = time.Now()
}
