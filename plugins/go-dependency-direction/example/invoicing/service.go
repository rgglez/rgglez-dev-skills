// Package invoicing turns an invoice into a message and sends it.
// It knows nothing about how mail is delivered.
package invoicing

import (
	"context"
	"fmt"
)

// sender is the whole of what this package asks of a mail transport: one
// method, written down here by the caller. No mail package is imported; the
// implementation is the side that has to fit.
type sender interface {
	Send(ctx context.Context, to, subject, body string) error
}

type Invoice struct {
	ID       string
	Customer string // email address
	Cents    int64
}

type Service struct {
	mail sender
}

// NewService accepts an interface, returns a struct.
func NewService(mail sender) *Service {
	return &Service{mail: mail}
}

func (s *Service) SendInvoice(ctx context.Context, inv Invoice) error {
	if inv.ID == "" {
		return fmt.Errorf("invoice: empty id")
	}
	if inv.Customer == "" {
		return fmt.Errorf("invoice %s: empty customer address", inv.ID)
	}
	if inv.Cents <= 0 {
		return fmt.Errorf("invoice %s: amount must be positive, got %d", inv.ID, inv.Cents)
	}

	subject := "Invoice " + inv.ID
	body := fmt.Sprintf("Amount due: %s", formatAmount(inv.Cents))

	if err := s.mail.Send(ctx, inv.Customer, subject, body); err != nil {
		return fmt.Errorf("invoice %s: send: %w", inv.ID, err)
	}
	return nil
}

// ponytail: USD only; take a currency argument when a second one shows up.
func formatAmount(cents int64) string {
	return fmt.Sprintf("$%d.%02d", cents/100, cents%100)
}
