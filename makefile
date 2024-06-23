include .env

up:
	docker-compose down
	docker-compose up --build -d
	docker image prune -f

down:
	docker-compose down

migration-create:
	~/go/bin/goose -dir ${DB_MIGRATION_DIR} ${DB_DRIVER} ${DB_DBSTRING} create $(name) sql

migration-up:
	~/go/bin/goose -dir ${DB_MIGRATION_DIR} ${DB_DRIVER} ${DB_DBSTRING} up

test:
	@go test -v ./internal/...

run:
	@go run ./cmd/api/main.go