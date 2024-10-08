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

type cardError string

const (
	NoError           cardError = ""
	InsufficientFunds cardError = "Insufficient funds"
	InvalidCard       cardError = "Invalid card number"
	ExpiredCard       cardError = "Card has expired"
	CVVMismatch       cardError = "CVV mismatch"
	CardBlocked       cardError = "Card is blocked"
	UnknownCardError  cardError = "Unknown card error"
)

func (bs *bankSimulator) SimulatePaymentProcessing(payment *models.Payment) (*models.BankResponse, error) {
	// Simulate processing delay
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	// Check for card errors
	err := bs.checkCard(payment.CardNumber, payment.CVV, payment.ExpirationDate)

	var approved bool
	var message string

	if err == NoError {
		approved = rand.Float32() < 0.7 // 70% approval rate
		if approved {
			message = "Payment approved"
		} else {
			message = "Payment declined by issuer"
		}
	} else {
		approved = false
		message = string(err)
	}

	response := &models.BankResponse{
		TransactionID: bs.generateTransactionID(),
		Approved:      approved,
		Message:       message,
	}

	return response, nil
}

func (bs *bankSimulator) checkCard(cardNumber, cvv string, expirationDate time.Time) cardError {
	errorChance := rand.Float32()
	switch {
	case errorChance < 0.02:
		return InsufficientFunds
	case errorChance < 0.04:
		return InvalidCard
	case errorChance < 0.06:
		return ExpiredCard
	case errorChance < 0.08:
		return CVVMismatch
	case errorChance < 0.09:
		return CardBlocked
	case errorChance < 0.10:
		return UnknownCardError
	default:
		return NoError
	}
}

func (bs *bankSimulator) generateTransactionID() string {
	return fmt.Sprintf("TR-%d", time.Now().UnixNano())
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
