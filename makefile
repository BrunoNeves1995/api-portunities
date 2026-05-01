.PHONY: default run build test docs clean lint fmt

APP_NAME	:= api-portunities-aula3
MAIN	:= main.go
DOCS_DIR	:= ./docs

default: run-with-docs

run:
	@CGO_ENABLED=1 go run $(MAIN)

run-with-docs:
	@swag init
	@CGO_ENABLED=1 go run $(MAIN)

build:
	@go build -o $(APP_NAME) $(MAIN)

test:
	@go test ./...

docs:
	@swag init

clean:
	@rm -rf $(APP_NAME) $(DOCS_DIR)

lint:
	@golangci-lint run

fmt:
	@go fmt ./...
