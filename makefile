DATABASE_URL = "postgres://postgres:postgres@localhost:7654/postgres?sslmode=disable"
DATABASE_MIGRATION_FILES = "file://scripts/database/migrations"
SOLUTION = "naive"
setup:
	make generate-sqlc; make docker-dev-env
run:
	DATABASE_URL=$(DATABASE_URL) DATABASE_MIGRATION_FILES=$(DATABASE_MIGRATION_FILES) go run .
docker-dev-env:
	docker compose up -d
cleanup:
	docker compose rm -s -v -f
generate-sqlc:
	sqlc generate