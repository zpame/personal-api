package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectToDatabase() (*pgx.Conn, error) {
	connectionString := os.Getenv("DATABASE_CONNECTION_STRING")

	if connectionString == "" {
		log.Fatal("DATABASE_CONNECTION_STRING environment variable is not set")
	}

	return pgx.Connect(context.Background(), connectionString)
}
