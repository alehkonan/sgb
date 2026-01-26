include .env
export

run-api:
	go tool air --build.cmd "go build -o ./tmp/main cmd/api/main.go"

run-bot:
	go tool air --build.cmd "go build -o ./tmp/main cmd/bot/main.go"

build-api:
	go build -o bin/api cmd/api/main.go

build-bot:
	go build -o bin/bot cmd/bot/main.go

seed-db:
	go run cmd/seed/main.go

.PHONY: run-api run-bot build-api build-bot seed-db
