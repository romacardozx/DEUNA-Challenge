package services

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
	"github.com/romacardozx/DEUNA-Challenge/internal/core/repositories"
	"github.com/romacardozx/DEUNA-Challenge/pkg/audit"
)

type RefundService interface {
	ProcessRefund(c *gin.Context, payload *models.RefundPayload) (*models.Refund, error)
	GetRefundDetails(c *gin.Context, refundID string) (*models.Refund, error)
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

func (s *refundService) ProcessRefund(c *gin.Context, payload *models.RefundPayload) (*models.Refund, error) {
	log := models.AuditData{}
	refund := models.Refund{
		ID:        uuid.New(),
		PaymentID: payload.PaymentID,
		Reason:    payload.Reason,
		CreatedAt: time.Now(),
	}
	payment, err := s.paymentRepository.GetByID(refund.PaymentID)
	if err != nil {
		return nil, err
	}
	refund.Amount = payment.Amount
	refund.Currency = payment.Currency

	bankResponse, err := s.bankSimulator.SimulateRefundProcessing(&refund)
	if err != nil {
		return nil, err
	}
	refund.TransactionID = bankResponse.TransactionID

	if bankResponse.Approved {
		refund.Status = "refunded"

		log.CompleteRefundLog(refund, *bankResponse, payment.MerchantID, payment.CustomerID)

		if err := audit.LogAudit(c, log); err != nil {
			return nil, err
		}

		err = s.refundRepository.Create(&refund)
		if err != nil {
			return nil, err
		}

		return &refund, nil

	}
	refund.Status = "declined"

	log.CompleteRefundLog(refund, *bankResponse, payment.MerchantID, payment.CustomerID)

	if err := audit.LogAudit(c, log); err != nil {
		return nil, err
	}

	return nil, errors.New("the refund was rejected")
}

func (s *refundService) GetRefundDetails(c *gin.Context, refundID string) (*models.Refund, error) {
	return s.refundRepository.GetByID(refundID)
}
