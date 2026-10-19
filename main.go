package main

import (
	"log"
	"net/http"

	"zack-personal-api/internal/controllers"
	"zack-personal-api/internal/daos"
	"zack-personal-api/internal/database"
)

func main() {
	conn, err := database.ConnectToDatabase()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	pizzaDAO := daos.NewPizzaDAO(conn)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /pizza-ratings", controllers.GetPizzaRatings(pizzaDAO))
	mux.HandleFunc("POST /pizza-ratings", controllers.CreatePizzaRating(pizzaDAO))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
