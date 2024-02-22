# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main


fmt: 
	@templ fmt .

generate: 
	templ generate

gen:
	+templ generate -watch -proxy=http://localhost:8080

air: 
	+air

play: fmt generate run
	

watch: 
	make --jobs=2 gen air

.PHONY: all build run test clean air generate fmt 
