package main

import (
	"log"

	"zack-personal-api/internal/database"
)

func main() {
	conn, err := database.ConnectToDatabase()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	log.Println("Successfully connected to the database.")
}
