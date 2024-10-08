package models

import "time"

type Merchant struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	Address         string    `json:"address"`
	BusinessType    string    `json:"business_type"`
	TaxID           string    `json:"tax_id"`
	BankAccountInfo string    `json:"bank_account_info"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
