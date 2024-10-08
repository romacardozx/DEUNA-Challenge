package models

type BankResponse struct {
	TransactionID string `json:"transaction_id"`
	Approved      bool   `json:"approved"`
	Message       string `json:"message"`
}
