---
name: go-dependency-direction
description: Applies dependency-direction rules when writing or reviewing Go code — consumer-defined minimal interfaces, no domain→infrastructure imports, one-way package arrows. Load whenever writing, refactoring, or reviewing Go packages, designing package layout, defining interfaces, or wiring repository/service/handler layers.
---

# Dependency direction in Go

The import graph **is** the architecture. Directory names (`domain/`, `usecase/`, `infrastructure/`) are decoration. A package called `domain` that imports the one holding your database pool is infrastructure code wearing a label.

Every layout decision reduces to one question: *which package is permitted to know that another exists?*

## Interfaces are declared by the caller

Put the interface in the package that **invokes** it, listing only the methods that package invokes. Declaring it beside the implementation forces every caller to import the provider, and any signature change there ripples outward to all of them.

Avoid — the provider exports the abstraction:

```go
// package mailer
type Mailer interface {
	Send(ctx context.Context, to, subject, body string) error
	SendBatch(ctx context.Context, to []string, subject, body string) error
	Close() error
}

type SMTPMailer struct{ /* ... */ }
func (m *SMTPMailer) Send(...) error { /* ... */ }
```

Prefer — the caller states the shape it depends on, and the implementation carries no project imports:

```go
// package invoicing
type sender interface {
	Send(ctx context.Context, to, subject, body string) error
}

type Service struct {
	mail sender
}

func NewService(mail sender) *Service {
	return &Service{mail: mail}
}
```

```go
// package smtpmail
type Client struct {
	addr string
	from string
	auth smtp.Auth
}

func (c *Client) Send(ctx context.Context, to, subject, body string) error {
	msg := fmt.Appendf(nil, "Subject: %s\r\n\r\n%s", subject, body)
	return smtp.SendMail(c.addr, c.auth, c.from, []string{to}, msg)
}
```

`smtpmail` has never heard of `invoicing`. `invoicing` has never heard of SMTP. Neither one has to change when the other does — that is what dependency inversion buys, and it comes from the import list, not from a picture.

Details worth keeping:

- **Keep the interface small.** One method here, because that is all `Service` calls — not the three `SMTPMailer` happens to publish.
- **Unexported is usually right.** `sender`, not `Sender`. No one outside needs to refer to it.
- **Accept interfaces, return structs.** Constructors take the caller's interface and hand back the concrete type.
- **This covers every collaborator**, not just storage: HTTP clients, notifiers, caches, clocks, metrics sinks, queue publishers.

## Two checks before you add an import

1. Which way does this new arrow run?
2. Is there already an arrow running the opposite way between these two packages?

And once per package: **strip the project down to this package alone — does it still build?** Packages holding types, arithmetic, validation, or wire formats must survive that.

## Layered designs are only as good as their arrows

`handler → service → repository`, with no return path. Two shortcuts erode it, both arriving as small conveniences:

- The service reaches for the handler's error-to-status helper. Fix: the service returns typed or sentinel errors, and the handler decides the status code.
- The repository borrows the service's DTO because the struct already exists. Fix: the repository returns its own model and the service converts.

Reject both. Writing a small struct twice costs less than an arrow pointing backwards.

## Check it, don't trust it

```sh
# every package's direct imports, one line each — the whole graph
go list -f '{{.ImportPath}} -> {{.Imports}}' ./...

# does anything reach a package it should not?
go list -deps ./... | grep <suspect>
```

Prefer `.Imports` over `.Deps`: `.Deps` is transitive and buries the answer under sixty lines of internal stdlib. Transport or persistence packages showing up in the core's import list is a genuine structural defect — and one that is invisible to any review that only looks at the diff.

## Cases this does not cover

- **`main` and wiring packages import freely.** Building the concrete `smtpmail.Client` and passing it to `NewService` is exactly their purpose.
- **Skip the interface when there is one implementation and no test fake.** The rule concerns direction, not universal abstraction.
- **Don't redeclare a standard-library interface.** When the shape you need already exists — `io.Writer`, `io.Reader`, `fmt.Stringer`, `sort.Interface`, `http.Handler` — take it instead of writing your own copy. It already satisfies the rule (one method, exported by no implementation), every reader knows it on sight, and it composes with everything built on it: `io.MultiWriter`, `io.Copy`, `bufio`, `httptest`. Write your own only for a shape the stdlib does not have, such as `Write` and `Flush` together.

---

Rules distilled from [*Forget clean architecture — Master dependency direction in Go*](https://blog.stackademic.com/forget-clean-architecture-master-dependency-direction-in-go-2c875193e940) by Yaninyz witty (Stackademic, July 2026). Wording and examples here are original.
