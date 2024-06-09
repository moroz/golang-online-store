guard-%:
	@ test -n "${$*}" || (echo "FATAL: Environment variable $* is not set!"; exit 1)

install:
	which goose || go install github.com/pressly/goose/v3/cmd/goose@latest
	which sqlc || go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	which templ || go install github.com/a-h/templ/cmd/templ@latest

db.create: guard-PGDATABASE
	createdb

db.migrate: guard-GOOSE_DBSTRING
	goose up

db.setup:
	make db.create
	make db.migrate
	make db.seed
	
db.reset: guard-PGDATABASE
	dropdb --force ${PGDATABASE}
	make db.setup

db.seed: guard-DATABASE_URL
	psql ${DATABASE_URL} -f db/seeds.sql

sqlc.gen:
	sqlc generate -f db/sqlc.yml

templ:
	templ generate

build:
	go build -o server -tags="sonic avx nomsgpack" .
