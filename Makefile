# ==============================================================================
# Main

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...
# ==============================================================================
# Docker compose commands
local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build

# ==============================================================================
# Tools commands

run-linter:
	echo "Starting linters"
	golangci-lint run ./...

swaggo:
	echo "Starting swagger generating"
	swag init -g **/**/*.go
