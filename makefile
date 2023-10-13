include .env

build:
	docker compose up -d

migrate-up:
	 migrate -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path ./db/migrations up

migrate-down:
	 migrate -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path ./db/migrations down

test :
	go test ./internal/...