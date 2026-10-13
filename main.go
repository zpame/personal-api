package main

import (
	"context"
	"log"
)

func main() {
	conn, err := connectToDatabase()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	log.Println("Successfully connected to the database.")
}
