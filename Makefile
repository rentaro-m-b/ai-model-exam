include .env
export $(shell sed 's/=.*//' .env)

MIGRATIONS_DIR=./migrations
SEEDS_DIR=./migrations/seed

# マイグレーションの適用
migrate-up:
	migrate -database "$(DATABASE_URL)" -path "$(MIGRATIONS_DIR)" up

# マイグレーションを一つ前に戻す
migrate-down:
	migrate -database "$(DATABASE_URL)" -path "$(MIGRATIONS_DIR)" down 1

# sqlc自動生成
generate:
	PGPASSWORD="${POSTGRES_PASSWORD}" pg_dump --schema-only --no-owner --no-privileges -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" -h localhost -p "${POSTGRES_PORT}" > db/schema.sql
	sqlc generate
