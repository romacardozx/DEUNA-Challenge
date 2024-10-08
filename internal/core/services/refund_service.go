package services

import (
	"errors"

	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/repositories"
)

type RefundService interface {
	ProcessRefund(refund *models.Refund) (*models.Refund, error)
	GetRefundDetails(refundID string) (*models.Refund, error)
	ListPaymentRefunds(paymentID string) ([]*models.Refund, error)
}

type refundService struct {
	refundRepository  repositories.RefundRepository
	paymentRepository repositories.PaymentRepository
	bankSimulator     BankSimulatorService
}

func NewRefundService(refundRepo repositories.RefundRepository, paymentRepo repositories.PaymentRepository, bankSim BankSimulatorService) RefundService {
	return &refundService{
		refundRepository:  refundRepo,
		paymentRepository: paymentRepo,
		bankSimulator:     bankSim,
	}
}

func (s *refundService) ProcessRefund(refund *models.Refund) (*models.Refund, error) {
	payment, err := s.paymentRepository.GetByID(refund.PaymentID)
	if err != nil {
		return nil, err
	}
	if payment.Status != "approved" {
		return nil, errors.New("payment cannot be refunded")
	}

	if refund.Amount > payment.Amount {
		return nil, errors.New("refund amount exceeds payment amount")
	}

	bankResponse, err := s.bankSimulator.SimulateRefundProcessing(refund)
	if err != nil {
		return nil, err
	}

	if bankResponse.Approved {
		refund.Status = "approved"
	} else {
		refund.Status = "declined"
	}
	refund.TransactionID = bankResponse.TransactionID

	err = s.refundRepository.Create(refund)
	if err != nil {
		return nil, err
	}

	return refund, nil
}

func (s *refundService) GetRefundDetails(refundID string) (*models.Refund, error) {
	return s.refundRepository.GetByID(refundID)
}

func (s *refundService) ListPaymentRefunds(paymentID string) ([]*models.Refund, error) {
	return s.refundRepository.ListByPayment(paymentID)
}
