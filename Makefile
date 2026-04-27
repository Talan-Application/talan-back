include .env
export

PROJECT_ROOT := $(CURDIR)

env-up:
	@docker compose up -d talan-postgres

env-down:
	@docker compose down

env-port-forwarder:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
ifndef name
	@echo Error: name is undefined.
	@echo Usage: make migrate-create name=init
	@exit 1
endif
	docker compose run --rm talan-postgres-migrate create -ext sql -dir /migrations -seq $(name)

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@docker compose run --rm talan-postgres-migrate \
		-path=/migrations/ \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@talan-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		$(action)

talan-run:
	@go run .\cmd\main.go