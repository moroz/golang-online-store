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
	make db.schema.dump
	
db.reset: guard-PGDATABASE
	dropdb --force ${PGDATABASE}
	createdb
	goose up

db.seed: guard-DATABASE_URL
	psql ${DATABASE_URL} -f db/seeds.sql

db.schema.dump: guard-PGDATABASE
	pg_dump -s -O -x --no-comments --no-publications ${PGDATABASE} -f db/schema.sql

sqlc.gen: db.schema.dump
	sqlc generate -f db/sqlc.yml

templ:
	templ generate

build:
	go build -o server -tags="sonic avx nomsgpack" .
