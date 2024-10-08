package services

import (
	"testing"

	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
)

func TestSimulatePaymentProcessing(t *testing.T) {
	simulator := NewBankSimulatorService()
	payment := &models.Payment{
		Amount:   100.00,
		Currency: "USD",
	}

	response, err := simulator.SimulatePaymentProcessing(payment)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if response.TransactionID == "" {
		t.Error("Expected a non-empty transaction ID")
	}

	if response.Message == "" {
		t.Error("Expected a non-empty message")
	}
}

func TestSimulateRefundProcessing(t *testing.T) {
	simulator := NewBankSimulatorService()
	refund := &models.Refund{
		Amount: 50.00,
	}

	response, err := simulator.SimulateRefundProcessing(refund)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if response.TransactionID == "" {
		t.Error("Expected a non-empty transaction ID")
	}

	if response.Message == "" {
		t.Error("Expected a non-empty message")
	}
}
