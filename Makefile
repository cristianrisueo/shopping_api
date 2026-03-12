include .env
export

build:
	go build -o bin/api cmd/api/main.go

run:
	go run cmd/api/main.go

lint:
	golangci-lint run

migrate-up:
	migrate -path db/migrations -database "$(DB_DSN)" up

migrate-down:
	migrate -path db/migrations -database "$(DB_DSN)" down

docker-up:
	docker compose -f docker/docker-compose.yml up -d

docker-down:
	docker compose -f docker/docker-compose.yml down