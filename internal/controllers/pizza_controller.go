package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

	"zack-personal-api/internal/dtos/requests"
	"zack-personal-api/internal/models"
	"zack-personal-api/internal/responses"
)

type pizzaRatingDAO interface {
	GetAllPizzaRatings() ([]models.PizzaRating, error)
}

type pizzaRatingCreator interface {
	CreatePizzaRating(ctx context.Context, rating models.PizzaRating) (models.PizzaRating, error)
}

func GetPizzaRatings(dao pizzaRatingDAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ratings, err := dao.GetAllPizzaRatings()
		if err != nil {
			http.Error(w, "failed to get pizza ratings", http.StatusInternalServerError)
			return
		}

		response := make([]responses.PizzaRatingResponse, 0, len(ratings))
		for _, rating := range ratings {
			response = append(response, toPizzaRatingResponse(rating))
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode pizza ratings", http.StatusInternalServerError)
		}
	}
}

func CreatePizzaRating(dao pizzaRatingCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if mediaType := r.Header.Get("Content-Type"); mediaType != "" &&
			!strings.EqualFold(strings.Split(mediaType, ";")[0], "application/json") {
			http.Error(w, "content type must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		var request requests.CreatePizzaRatingRequest
		if err := decoder.Decode(&request); err != nil {
			var maxBytesError *http.MaxBytesError
			if errors.As(err, &maxBytesError) {
				http.Error(w, "request body is too large", http.StatusRequestEntityTooLarge)
				return
			}
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		if err := decoder.Decode(&struct{}{}); err != io.EOF {
			http.Error(w, "request body must contain a single JSON object", http.StatusBadRequest)
			return
		}

		rating, err := pizzaRatingFromRequest(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		createdRating, err := dao.CreatePizzaRating(r.Context(), rating)
		if err != nil {
			http.Error(w, "failed to create pizza rating", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toPizzaRatingResponse(createdRating)); err != nil {
			http.Error(w, "failed to encode pizza rating", http.StatusInternalServerError)
		}
	}
}

func pizzaRatingFromRequest(request requests.CreatePizzaRatingRequest) (models.PizzaRating, error) {
	if request.Restaurant == nil || request.Location == nil || request.DateVisited == nil ||
		request.PizzaOrdered == nil || request.Rating == nil || request.Recommend == nil {
		return models.PizzaRating{}, errors.New("restaurant, location, date_visited, pizza_ordered, rating, and recommend are required")
	}

	for _, field := range []struct {
		name  string
		value string
	}{
		{name: "restaurant", value: *request.Restaurant},
		{name: "location", value: *request.Location},
		{name: "pizza_ordered", value: *request.PizzaOrdered},
	} {
		if strings.TrimSpace(field.value) == "" {
			return models.PizzaRating{}, fmt.Errorf("%s must not be empty", field.name)
		}
		if utf8.RuneCountInString(field.value) > 100 {
			return models.PizzaRating{}, fmt.Errorf("%s must be at most 100 characters", field.name)
		}
	}

	dateVisited, err := time.Parse("2006-01-02", *request.DateVisited)
	if err != nil {
		return models.PizzaRating{}, errors.New("date_visited must be a valid date in YYYY-MM-DD format")
	}
	if math.Abs(*request.Rating) >= 99.995 {
		return models.PizzaRating{}, errors.New("rating must be between -99.99 and 99.99")
	}

	return models.PizzaRating{
		Restaurant:   *request.Restaurant,
		Location:     *request.Location,
		DateVisited:  dateVisited,
		PizzaOrdered: *request.PizzaOrdered,
		Rating:       *request.Rating,
		Reccomended:  *request.Recommend,
	}, nil
}

func toPizzaRatingResponse(rating models.PizzaRating) responses.PizzaRatingResponse {
	return responses.PizzaRatingResponse{
		ID:           rating.ID,
		Restaurant:   rating.Restaurant,
		Location:     rating.Location,
		PizzaOrdered: rating.PizzaOrdered,
		Rating:       rating.Rating,
		Reccomend:    rating.Reccomended,
	}
}
