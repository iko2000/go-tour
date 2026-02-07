.PHONY: build run test lint migrate

build:
	go build -o bin/server cmd/server/main.go

run:
	go run cmd/server/main.go		
		
test:
	go test ./... -v

lint:
	golangci-lint run

migrate:
	go run cmd/migrate/main.go