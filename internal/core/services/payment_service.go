package services

import (
	"time"

	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/repositories"
)

type PaymentService interface {
	ProcessPayment(payment *models.Payment) (*models.Payment, error)
	GetPaymentDetails(paymentID string) (*models.Payment, error)
	ListMerchantPayments(merchantID string, limit, offset int) ([]*models.Payment, error)
}

type paymentService struct {
	paymentRepo   repositories.PaymentRepository
	bankSimulator BankSimulatorService
}

func NewPaymentService(paymentRepo repositories.PaymentRepository, bankSimulator BankSimulatorService) PaymentService {
	return &paymentService{
		paymentRepo:   paymentRepo,
		bankSimulator: bankSimulator,
	}
}

func (s *paymentService) ProcessPayment(payment *models.Payment) (*models.Payment, error) {
	bankResponse, err := s.bankSimulator.SimulatePaymentProcessing(payment)
	if err != nil {
		return nil, err
	}
	if bankResponse.Approved {
		payment.Status = "approved"
	} else {
		payment.Status = "declined"
	}
	payment.TransactionID = bankResponse.TransactionID
	payment.CreatedAt = time.Now()

	err = s.paymentRepo.Create(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *paymentService) GetPaymentDetails(paymentID string) (*models.Payment, error) {
	return s.paymentRepo.GetByID(paymentID)
}

func (s *paymentService) ListMerchantPayments(merchantID string, limit, offset int) ([]*models.Payment, error) {
	return s.paymentRepo.ListByMerchant(merchantID, limit, offset)
}
