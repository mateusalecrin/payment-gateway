package dto

import (
	"time"

	"github.com/mateusalecrin/payment-gateway/go-api/internal/domain"
)

const (
	StatusPending  = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

type CreateInvoiceInput struct {
	APIKey          string
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	PaymentType     string  `json:"payment_type"`
	CardNumber      string  `json:"card_number"`
	HolderName      string  `json:"holder_name"`
	ExpirationMonth int     `json:"expiration_month"`
	ExpirationYear  int     `json:"expiration_year"`
	CVV             string  `json:"cvv"`
}

type InvoiceOutput struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToInvoice(input CreateInvoiceInput, accountID string) (*domain.Invoice, error) {
	card := domain.CreditCard{
		Number:          input.CardNumber,
		HolderName:      input.HolderName,
		ExpirationMonth: input.ExpirationMonth,
		ExpirationYear:  input.ExpirationYear,
		CVV:             input.CVV,
	}

	return domain.NewInvoice(
		accountID,
		input.Amount,
		input.Description,
		input.PaymentType,
		card,
	)
}

func FromInvoice(invoice *domain.Invoice) *InvoiceOutput {
	return &InvoiceOutput{
		ID:             invoice.ID,
		AccountID:      invoice.AccountID,
		Amount:         invoice.Amount,
		Status:         string(invoice.Status),
		Description:    invoice.Description,
		PaymentType:    invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt:      invoice.CreatedAt,
		UpdatedAt:      invoice.UpdatedAt,
	}
}
