include .env
export

dev:
	go tool air

build:
	go build -o bot cmd/bot/main.go

.PHONY: dev
