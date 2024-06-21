.PHONY: migrate dc_build dc_up dc_down logs deps

migrate:
	go run internal/migrations/migrate.go

dc_build:
	docker compose -f ./.docker/docker-compose.yml up --build

dc_up:
	docker compose -f ./.docker/docker-compose.yml up -d

dc_down:
	docker compose -f ./.docker/docker-compose.yml down

logs:
	docker compose -f ./.docker/docker-compose.yml logs -f

deps:
	go mod tidy
