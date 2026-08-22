// Composition root: the only place that imports both sides and names the
// concrete implementation.
package main

import (
	"context"
	"fmt"
	"log"

	"example/billing/invoicing"
	"example/billing/memmail"
)

func main() {
	outbox := memmail.New()
	svc := invoicing.NewService(outbox)

	inv := invoicing.Invoice{ID: "INV-1", Customer: "ada@example.com", Cents: 4999}
	if err := svc.SendInvoice(context.Background(), inv); err != nil {
		log.Fatal(err)
	}

	for _, m := range outbox.Sent() {
		fmt.Printf("to=%s subject=%q body=%q\n", m.To, m.Subject, m.Body)
	}
}
