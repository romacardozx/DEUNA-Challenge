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

type PaymentService interface {
	ProcessPayment(c *gin.Context, paymentPayload *models.PaymentPayload) (*models.Payment, error)
	GetPaymentDetails(c *gin.Context, paymentID string) (*models.Payment, error)
	ListMerchantPayments(c *gin.Context, merchantID string, limit, offset int) ([]models.Payment, error)
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

func (s *paymentService) ProcessPayment(c *gin.Context, paymentPayload *models.PaymentPayload) (*models.Payment, error) {
	log := models.AuditData{}
	payment := models.Payment{
		CustomerID:  paymentPayload.CustomerID,
		MerchantID:  paymentPayload.MerchantID,
		Amount:      paymentPayload.Amount,
		Currency:    paymentPayload.Currency,
		Description: paymentPayload.Description,
		CreatedAt:   time.Now(),
	}
	payment.ID = uuid.New()
	bankResponse, err := s.bankSimulator.SimulatePaymentProcessing(&payment)
	if err != nil {
		return nil, err
	}

	payment.TransactionID = bankResponse.TransactionID
	if bankResponse.Approved {
		payment.Status = "approved"

		log.CompletePaymentLog(payment, *bankResponse)

		if err := audit.LogAudit(c, log); err != nil {
			return nil, err
		}

		err = s.paymentRepo.Create(&payment)
		if err != nil {
			return nil, err
		}

		return &payment, nil

	}
	payment.Status = "rejected"

	log.CompletePaymentLog(payment, *bankResponse)

	if err := audit.LogAudit(c, log); err != nil {
		return nil, err
	}

	return nil, errors.New("the payment was rejected")

}

func (s *paymentService) GetPaymentDetails(c *gin.Context, paymentID string) (*models.Payment, error) {
	return s.paymentRepo.GetByID(paymentID)
}

func (s *paymentService) ListMerchantPayments(c *gin.Context, merchantID string, limit, offset int) ([]models.Payment, error) {
	return s.paymentRepo.ListByMerchant(merchantID, limit, offset)
}
