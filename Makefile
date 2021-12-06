postgres:
	sudo docker run --name postgres-latest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=simple_bank -d postgres:latest

createdb:
	sudo docker exec -it postgres-latest createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres-latest dropdb simple_bank

migrateup:
	migrate -path db/migration/ -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc