package main

import (
	"context"
	"log"
	"zack-personal-api/internal/database"
)

func main() {
	conn, err := database.ConnectToDatabase()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	log.Println("Successfully connected to the database.")
}
