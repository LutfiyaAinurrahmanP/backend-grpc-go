# Create migration
migrate create -ext sql -dir db/migration -seq init_schema

#  Run Cli postgresql
docker exec -it postgres12 /bin/sh
docker exec -it postgres12 psql -U root