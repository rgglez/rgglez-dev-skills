# Go dependency direction

The import graph **is** the architecture. Directory names (`domain/`, `usecase/`, `infrastructure/`) are decoration — a package called `domain` that imports the one holding your database pool is infrastructure code wearing a label.

## Interfaces are declared by the caller

Put the interface in the package that **invokes** it, listing only the methods that package invokes. Declaring it beside the implementation forces every caller to import the provider, and a signature change there ripples to all of them.

```go
// package invoicing — the caller states the shape it depends on
type sender interface {
	Send(ctx context.Context, to, subject, body string) error
}

func NewService(mail sender) *Service { return &Service{mail: mail} }
```

```go
// package smtpmail — satisfies it incidentally, imports nothing of ours
func (c *Client) Send(ctx context.Context, to, subject, body string) error { /* ... */ }
```

- Small interface: one method, not the three the implementation publishes.
- Unexported is usually right (`sender`, not `Sender`).
- Accept interfaces, return structs.
- Covers every collaborator: storage, HTTP clients, notifiers, caches, clocks, metrics sinks, queues.

## Two checks before you add an import

1. Which way does this new arrow run?
2. Is there already an arrow running the opposite way between these two packages?

Once per package: **strip the project down to this package alone — does it still build?** Packages holding types, validation, or wire formats must survive that.

## Layered designs

`handler → service → repository`, no return path. Two shortcuts erode it:

- Service reaching for the handler's error-to-status helper → return typed errors instead; the handler picks the status code.
- Repository borrowing the service's DTO → the repository returns its own model, the service converts.

Writing a small struct twice costs less than an arrow pointing backwards.

## Check it

```sh
# every package's direct imports, one line each — the whole graph
go list -f '{{.ImportPath}} -> {{.Imports}}' ./...

# does anything reach a package it should not?
go list -deps ./... | grep <suspect>
```

Prefer `.Imports` over the transitive `.Deps`, which buries the answer under internal stdlib. Transport or persistence packages in the core's import list is a structural defect, invisible to any diff-only review.

## Cases this does not cover

- `main` and wiring packages import freely — that is their purpose.
- No interface for one implementation with no test fake.
- Don't redeclare a stdlib interface. If `io.Writer`, `fmt.Stringer`, `sort.Interface` or the like already matches the shape you need, take it: already caller-shaped, instantly recognizable, and it composes with `io.Copy`, `bufio`, `httptest`. Declare your own only for a shape the stdlib lacks (`Write` + `Flush` together).

---

Rules distilled from [*Forget clean architecture — Master dependency direction in Go*](https://blog.stackademic.com/forget-clean-architecture-master-dependency-direction-in-go-2c875193e940) by Yaninyz witty (Stackademic, July 2026). Wording and examples here are original.
