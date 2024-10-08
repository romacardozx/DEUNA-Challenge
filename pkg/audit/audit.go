package audit

import (
	"context"

	"github.com/romacardozx/DEUNA-Challenge/internal/core/models"
	"github.com/romacardozx/DEUNA-Challenge/internal/database"
)

func LogAudit(ctx context.Context, data models.AuditData) error {
	db := database.GetDB()
	_, err := db.ExecContext(ctx,
		`SELECT insert_audit_log($1, $2, $3, $4, $5, $6, $7, $8)`,
		data.PaymentID,
		data.Amount,
		data.MerchantID,
		data.CustomerID,
		data.Currency,
		data.TransactionID,
		data.Status,
		data.Message,
	)
	return err
}
