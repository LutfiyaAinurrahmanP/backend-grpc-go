postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:12-alpine

start:
	docker container start postgres12

stop :
	docker container stop postgres12

cli:
	docker exec -it postgres12 psql -U root 

clidb:
	docker exec -it postgres12 psql -U root -d bank_app

createdb:
	docker exec -it postgres12 createdb -U root bank_app

dropdb:
	docker exec -it postgres12 dropdb -U root bank_app

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank_app?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

.PHONY: createdb dropdb migrateup migratedown start stop cli clidb postgres sqlc test