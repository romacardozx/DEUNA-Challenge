package services

import (
	"fmt"

	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/repositories"
)

func CreateRefund(paymentID string, amount float64, reason string) (*models.Refund, error) {
	payment, err := repositories.GetPayment(paymentID)
	if err != nil {
		return nil, err
	}

	if payment.Amount < amount {
		return nil, fmt.Errorf("refund amount exceeds payment amount")
	}

	refund := &models.Refund{
		PaymentID: paymentID,
		Amount:    amount,
		Reason:    reason,
		Status:    "pending",
	}

	createdRefund, err := repositories.CreateRefund(refund)
	if err != nil {
		return nil, err
	}

	return createdRefund, nil
}

func GetRefund(id string) (*models.Refund, error) {
	return repositories.GetRefund(id)
}
