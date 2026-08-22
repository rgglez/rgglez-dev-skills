package invoicing

import (
	"context"
	"errors"
	"testing"
)

// The payoff of a one-method caller-side interface: the fake is four lines
// and the test imports no mail package at all.
type fakeSender struct {
	to, subject, body string
	err               error
}

func (f *fakeSender) Send(_ context.Context, to, subject, body string) error {
	if f.err != nil {
		return f.err
	}
	f.to, f.subject, f.body = to, subject, body
	return nil
}

func TestSendInvoice(t *testing.T) {
	ctx := context.Background()

	mail := &fakeSender{}
	inv := Invoice{ID: "INV-1", Customer: "ada@example.com", Cents: 4999}
	if err := NewService(mail).SendInvoice(ctx, inv); err != nil {
		t.Fatalf("valid invoice: %v", err)
	}
	if mail.to != "ada@example.com" || mail.subject != "Invoice INV-1" {
		t.Errorf("to=%q subject=%q", mail.to, mail.subject)
	}
	if mail.body != "Amount due: $49.99" {
		t.Errorf("body = %q", mail.body)
	}

	bad := []Invoice{
		{ID: "", Customer: "ada@example.com", Cents: 1},
		{ID: "INV-2", Customer: "", Cents: 1},
		{ID: "INV-3", Customer: "ada@example.com", Cents: 0},
		{ID: "INV-4", Customer: "ada@example.com", Cents: -5},
	}
	for _, inv := range bad {
		if err := NewService(&fakeSender{}).SendInvoice(ctx, inv); err == nil {
			t.Errorf("SendInvoice(%+v) = nil, want error", inv)
		}
	}

	boom := errors.New("smtp unreachable")
	err := NewService(&fakeSender{err: boom}).SendInvoice(ctx, inv)
	if !errors.Is(err, boom) {
		t.Errorf("transport failure not wrapped: %v", err)
	}
}

func TestFormatAmount(t *testing.T) {
	for _, c := range []struct {
		cents int64
		want  string
	}{{4999, "$49.99"}, {5, "$0.05"}, {100, "$1.00"}, {1234567, "$12345.67"}} {
		if got := formatAmount(c.cents); got != c.want {
			t.Errorf("formatAmount(%d) = %s, want %s", c.cents, got, c.want)
		}
	}
}
