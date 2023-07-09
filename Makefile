postgresInit:
	docker run --name pgsql-dev -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Welcome4 -d postgres:15-alpine

postgres:
	docker exec -it pgsql-dev psql

createdb:
	docker exec -it pgsql-dev createdb --username=root --owner=postgres postgres

dropdb:
	docker exec -it pgsql-dev dropdb postgres

migrateup:
	migrate -path connector/db/migrations -database "postgresql://root:Welcome4@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path connector/db/migrations -database "postgresql://root:Welcome4@localhost:5432/postgres?sslmode=disable" -verbose down

.PHONY: postgresInit postgres createdb dropdb migrateup migratedown
