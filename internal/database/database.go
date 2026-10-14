package database

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDatabase() (*pgxpool.Pool, error) {
	connectionString := os.Getenv("DATABASE_CONNECTION_STRING")

	if connectionString == "" {
		return nil, errors.New("DATABASE_CONNECTION_STRING is not set")
	}

	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 4

	pool, err := pgxpool.NewWithConfig(
		context.Background(),
		config,
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
