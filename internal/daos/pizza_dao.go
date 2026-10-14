package daos

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"zack-personal-api/internal/models"
)

var errGettingPizzaRatings = errors.New("failed to get pizza ratings")

type PizzaDAO struct {
	db *pgxpool.Pool
}

func NewPizzaDAO(db *pgxpool.Pool) *PizzaDAO {
	return &PizzaDAO{db: db}
}

func (dao *PizzaDAO) GetAllPizzaRatings() ([]models.PizzaRating, error) {
	query, err := os.ReadFile("sql/migrations/GetPizzaRatings.sql")
	if err != nil {
		return nil, errGettingPizzaRatings
	}

	rows, err := dao.db.Query(context.Background(), string(query))
	if err != nil {
		return nil, errGettingPizzaRatings
	}
	defer rows.Close()

	var ratings []models.PizzaRating
	return ratings, errGettingPizzaRatings // Placeholder for actual implementation
}
