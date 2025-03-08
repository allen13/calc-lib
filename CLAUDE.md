# CLAUDE.md - AI Assistant Guide

## Build & Test Commands
- Build: `go build ./...`
- Run: `go run main.go`
- Format: `go fmt ./...`
- Lint: `golint ./...`
- Vet: `go vet ./...`
- Test all: `go test ./...`
- Test single: `go test -run "TestName" ./package/path`
- Test with coverage: `go test -cover ./...`

## Code Style Guidelines
- Follow Go's official style guide and conventions
- Format with `gofmt` or `go fmt`
- Use `golint` and `go vet` for code quality
- Organize imports alphabetically (standard library first, then third-party)
- Naming: use camelCase for unexported and PascalCase for exported identifiers
- Error handling: check errors explicitly, avoid panic in production code
- Prefer composition over inheritance
- Use interfaces for abstraction and testability
- Document public APIs with godoc-style comments
- Write table-driven tests for all business logic
- Follow the principle of "Accept interfaces, return structs"