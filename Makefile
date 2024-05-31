guard-%:
	@ test -n "${$*}" || (echo "FATAL: Environment variable $* is not set!"; exit 1)

install:
	which goose || go install github.com/pressly/goose/v3/cmd/goose@latest

db.create: guard-PGDATABASE
	createdb

db.migrate:
	goose up
	
db.reset: guard-PGDATABASE
	dropdb --force ${PGDATABASE}
	createdb
	goose up
