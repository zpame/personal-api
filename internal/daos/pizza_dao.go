package daos

import (
	"zack-personal-api/internal/responses"
)

type PizzaDAO interface {
	GetPizzaRatings() ([]responses.PizzaRatingResponse, error)
}
