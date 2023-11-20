fmt:
	docker compose exec api go fmt ./...
api:
	docker compose exec api sh
migrateup:
	docker compose exec api go run cmd/migrate_up/main.go
migratedown:
	docker compose exec api go run cmd/migrate_down/main.go
tidy:
	docker compose exec api go mod tidy