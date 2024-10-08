package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
)

type BankSimulatorService interface {
	SimulatePaymentProcessing(payment *models.Payment) (*models.BankResponse, error)
	SimulateRefundProcessing(refund *models.Refund) (*models.BankResponse, error)
}

type bankSimulator struct{}

func NewBankSimulatorService() BankSimulatorService {
	return &bankSimulator{}
}

func (bs *bankSimulator) SimulatePaymentProcessing(payment *models.Payment) (*models.BankResponse, error) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	approved := rand.Float32() < 0.5

	response := &models.BankResponse{
		TransactionID: generateTransactionID(),
		Approved:      approved,
		Message:       getBankMessage(approved),
	}

	return response, nil
}

func (bs *bankSimulator) SimulateRefundProcessing(refund *models.Refund) (*models.BankResponse, error) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	approved := rand.Float32() < 0.8

	response := &models.BankResponse{
		TransactionID: generateTransactionID(),
		Approved:      approved,
		Message:       getBankMessage(approved),
	}

	return response, nil
}

func generateTransactionID() string {
	return fmt.Sprintf("TR-%d", time.Now().UnixNano())
}

func getBankMessage(approved bool) string {
	if approved {
		return "Transaction approved"
	}
	return "Transaction declined"
}
