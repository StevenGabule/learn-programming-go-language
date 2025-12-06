# Go Learning Roadmap for JavaScript Developers

Here's a structured outline tailored to your JS background, John Paul. I've organized it to leverage what you already know while highlighting where Go differs significantly.

---

## Phase 1: Foundations (1-2 weeks)

**Environment Setup**

- Install Go, set up GOPATH and Go modules
- VS Code with Go extension (or GoLand)
- Understand `go mod init`, `go run`, `go build`

**Syntax & Basic Types**

- Variables: `var` vs `:=` (short declaration)
- Basic types: `int`, `float64`, `string`, `bool`
- Zero values (unlike JS's `undefined`)
- Constants with `const` and `iota`
- No implicit type coercion (big shift from JS)

**Control Structures**

- `if`, `for` (Go's only loop — no `while`)
- `switch` (much more powerful than JS, no fall-through by default)
- `defer` — unique to Go, crucial for resource cleanup

**Functions**

- Multiple return values (very common pattern)
- Named return values
- Variadic functions
- Functions as first-class citizens (closures work similarly to JS)

---

## Phase 2: Core Go Concepts (2-3 weeks)

**Composite Types**

- Arrays (fixed size) vs Slices (dynamic, what you'll mostly use)
- Maps (similar to JS objects/Maps)
- Structs (closest thing to classes, but no inheritance)

**Pointers**

- Value vs reference semantics (JS hides this from you)
- `&` (address-of) and `*` (dereference)
- Why Go uses pointers explicitly vs JS's hidden references

**Methods & Interfaces**

- Methods on structs (receiver functions)
- Interfaces: implicit satisfaction (no `implements` keyword)
- Empty interface `interface{}` / `any` (like JS's `any` in TS)
- Interface composition

**Error Handling**

- No exceptions — explicit `error` returns
- The `if err != nil` pattern (you'll write this a lot)
- Custom error types
- `errors.Is()`, `errors.As()`, wrapping errors

---

## Phase 3: Concurrency (2-3 weeks)

This is Go's killer feature and likely new territory coming from JS's single-threaded model.

**Goroutines**

- Lightweight threads with `go` keyword
- Contrast with JS Promises/async-await
- How the Go scheduler works

**Channels**

- Communication between goroutines
- Buffered vs unbuffered channels
- `select` statement for multiplexing
- Channel directions (send-only, receive-only)

**Concurrency Patterns**

- Worker pools
- Fan-in, fan-out
- Context for cancellation and timeouts
- `sync` package: `WaitGroup`, `Mutex`, `Once`

---

## Phase 4: Standard Library Deep Dive (2 weeks)

Go's stdlib is excellent — you'll reach for external packages far less than in Node.

**Essential Packages**

- `fmt` — formatting and printing
- `strings`, `strconv` — string manipulation
- `io`, `bufio` — I/O operations
- `os`, `path/filepath` — file system operations
- `encoding/json` — JSON marshaling/unmarshaling (struct tags)
- `net/http` — HTTP client and server (production-ready out of the box)
- `time` — time handling (very different from JS Date)
- `context` — request-scoped values, cancellation
- `testing` — built-in testing framework

---

## Phase 5: Project Structure & Tooling (1 week)

**Go Modules**

- `go.mod` and `go.sum`
- Versioning and dependency management
- Private modules

**Project Layout**

- Standard Go project structure
- `cmd/`, `internal/`, `pkg/` conventions
- When to split into packages

**Tooling**

- `go fmt` — code formatting (non-negotiable in Go)
- `go vet` — static analysis
- `go test` — testing with coverage
- `golangci-lint` — comprehensive linting
- `go generate` — code generation

---

## Phase 6: Web Development & APIs (2-3 weeks)

Coming from Node/Express, this will feel familiar conceptually.

**Standard Library HTTP**

- `http.Handler` interface
- Building REST APIs with just stdlib
- Middleware patterns

**Popular Frameworks & Routers**

- **Gin** — most popular, Express-like feel
- **Echo** — similar to Gin, slightly different API
- **Chi** — lightweight, stdlib-compatible
- **Fiber** — Express-inspired, built on fasthttp

**Database Access**

- `database/sql` — standard interface
- **sqlx** — extensions to database/sql
- **GORM** — ORM (similar to Sequelize/TypeORM)
- **sqlc** — generates type-safe code from SQL
- **pgx** — PostgreSQL driver

**Other Web Essentials**

- **validator** — struct validation
- **jwt-go** — JWT handling
- **viper** — configuration management
- **zap** or **zerolog** — structured logging

---

## Phase 7: Advanced Topics (Ongoing)

**Generics (Go 1.18+)**

- Type parameters
- Constraints
- Generic data structures

**Reflection**

- `reflect` package
- When to use (rarely) and when to avoid

**Performance**

- Benchmarking with `go test -bench`
- Profiling with `pprof`
- Memory management and escape analysis

**Building CLI Tools**

- **Cobra** — CLI framework (used by Docker, Kubernetes)
- **Bubble Tea** — terminal UI framework

---

## Phase 8: Ecosystem & Production (Ongoing)

**Microservices & gRPC**

- Protocol Buffers
- gRPC services
- **Connect** — modern RPC framework

**DevOps Integration**

- Building minimal Docker images with multi-stage builds
- Cross-compilation (`GOOS`, `GOARCH`)
- Kubernetes operators (if relevant)

**Observability**

- OpenTelemetry for tracing
- Prometheus metrics
- Structured logging patterns

---

## Key Mindset Shifts from JS

| JavaScript                   | Go                                |
| ---------------------------- | --------------------------------- |
| Dynamic typing               | Static typing                     |
| Exceptions                   | Explicit error returns            |
| Single-threaded + event loop | Goroutines + channels             |
| Classes & inheritance        | Structs & composition             |
| npm/yarn ecosystem           | Smaller, stdlib-focused ecosystem |
| Implicit `this`              | Explicit receivers                |
| Truthy/falsy                 | Strict boolean                    |

---

## Recommended Resources

**Official**

- [Go Tour](https://go.dev/tour) — interactive intro
- [Effective Go](https://go.dev/doc/effective_go) — idiomatic patterns
- [Go by Example](https://gobyexample.com) — practical snippets

**Books**

- "Learning Go" by Jon Bodner (modern, practical)
- "Concurrency in Go" by Katherine Cox-Buday

**Practice**

- [Exercism Go Track](https://exercism.org/tracks/go)
- Build a REST API with Gin + PostgreSQL
- Build a CLI tool with Cobra

---

Want me to dive deeper into any section, or create a starter project outline to get you hands-on quickly?
