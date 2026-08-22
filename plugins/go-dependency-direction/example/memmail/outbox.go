// Package memmail collects messages in memory instead of delivering them.
// It imports nothing from this project: delete every other package here and
// it still compiles.
package memmail

import (
	"context"
	"sync"
)

type Message struct {
	To      string
	Subject string
	Body    string
}

type Outbox struct {
	mu   sync.Mutex
	sent []Message
}

func New() *Outbox { return &Outbox{} }

// ponytail: one mutex for the whole slice; nothing here is hot enough to shard.
func (o *Outbox) Send(_ context.Context, to, subject, body string) error {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.sent = append(o.sent, Message{To: to, Subject: subject, Body: body})
	return nil
}

func (o *Outbox) Sent() []Message {
	o.mu.Lock()
	defer o.mu.Unlock()
	return append([]Message(nil), o.sent...)
}
