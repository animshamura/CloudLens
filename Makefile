.PHONY: test build run docker-up docker-down migrate

all: test

test:
	go test ./...

build:
	go build -o api-gateway ./cmd/api-gateway

run: build
	./api-gateway

migrate:
	bash scripts/run-migrations.sh

docker-up:
	docker compose up --build -d

docker-down:
	docker compose down -v
