package main

import (
	"database/sql"
	"log"

	"github.com/LutfiyaAinurrahmanP/backend-grpc-go/api"
	db "github.com/LutfiyaAinurrahmanP/backend-grpc-go/db/sqlc"
	"github.com/LutfiyaAinurrahmanP/backend-grpc-go/utils"
	_ "github.com/lib/pq" // Import driver PostgreSQL
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
