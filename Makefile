DB_URL=postgres://postgres:root@localhost:5432/game_shop?sslmode=disable
MIGRATION_PATH=migrations

.PHONY: migrate-up migrate-down migrate-create migrate-force migrate-version help

migrate-up:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" up

	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" down 1

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

migrate-force:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" force $(version)

migrate-version:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" version

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'