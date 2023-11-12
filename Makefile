fmt:
	docker compose exec api go fmt ./...
api:
	docker compose exec api sh