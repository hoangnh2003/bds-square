GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING = "root:root1234@tcp(127.0.0.1:33306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema
CONTAINER_ID := $(shell docker ps --filter "name=app" --format "{{.ID}}" | head -n 1)

run:
	docker-compose up -d

debug:
	@echo "🔍 Attaching debugger to container $(CONTAINER_ID)..."
	@PID=$$(docker exec $(CONTAINER_ID) pgrep -f "/app/app"); \
	if [ -z "$$PID" ]; then \
		echo "❌ Không tìm thấy tiến trình Go (/app/app) để attach"; \
		docker exec $(CONTAINER_ID) ps aux | grep app; \
		exit 1; \
	fi; \
	echo "✅ Tìm thấy PID $$PID. Attaching..."; \
	docker exec -it $(CONTAINER_ID) dlv attach $$PID --headless --listen=:2345 --api-version=2 --accept-multiclient

stop-debug:
	@docker exec -it $(CONTAINER_ID) pkill -f 'dlv attach' || true

migrate:
	docker-compose exec app go run ./cmd/migrate/main.go

db-up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

db-down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

db-reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

.PHONY: db-up db-down db-reset
