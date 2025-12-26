package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/LutfiyaAinurrahmanP/backend-grpc-go/utils"
	_ "github.com/lib/pq" // Import driver PostgreSQL
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
