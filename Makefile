.PHONY: format
format:
	@go fmt .

.PHONY: lint
lint:
	@go vet
	@golint ./cmd/...

.PHONY: test
test:
	@go test ./...
