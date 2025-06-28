.PHONY: dev build run test clean

# Development with live reload
dev:
	$(shell go env GOPATH)/bin/air

# Build the application
build:
	go build -o bin/app main.go

# Run the application (production)
run:
	go run main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf tmp/
	rm -rf bin/
	rm -f build-errors.log

# Install dependencies
deps:
	go mod tidy
	go mod download

# Install Air for live reload
install-air:
	go install github.com/air-verse/air@latest 