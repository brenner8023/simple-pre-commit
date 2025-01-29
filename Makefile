
install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go mod tidy

test:
	go test -race -v ./tests

coverage-ci:
	go test -v -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html

coverage:
	@make coverage-ci
	open coverage.html

pre-commit:
	@make lint

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...

lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...

build:
	go build -o simple-pre-commit
