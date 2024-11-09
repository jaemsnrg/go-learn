# Packages

## Golang

### Migrate - Assists with database migrations

https://github.com/golang-migrate/migrate

brew install golang-migrate


### commands -

# Create a new migration file
# migrate create: Creates a new migration file
# -ext sql: Sets the file extension to .sql
# -dir db/migrations: Places migration files in db/migrations directory
# -seq: Uses sequential version numbers (e.g. 000001, 000002)
# init_schema: Name of the migration

migrate create -ext sql -dir db/migrations -seq init_schema

