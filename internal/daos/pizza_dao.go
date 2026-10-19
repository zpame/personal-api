package daos

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"zack-personal-api/internal/models"
)

var errGettingPizzaRatings = errors.New("failed to get pizza ratings")
var errCreatingPizzaRating = errors.New("failed to create pizza rating")

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
	for rows.Next() {
		var rating models.PizzaRating
		if err := rows.Scan(
			&rating.ID,
			&rating.Restaurant,
			&rating.Location,
			&rating.DateVisited,
			&rating.PizzaOrdered,
			&rating.Rating,
			&rating.Reccomended,
			&rating.CreatedAt,
		); err != nil {
			return nil, errGettingPizzaRatings
		}
		ratings = append(ratings, rating)
	}

	if rows.Err() != nil {
		return nil, errGettingPizzaRatings
	}

	return ratings, nil
}

func (dao *PizzaDAO) CreatePizzaRating(ctx context.Context, rating models.PizzaRating) (models.PizzaRating, error) {
	query, err := os.ReadFile("sql/migrations/InsertPizzaRating.sql")
	if err != nil {
		return models.PizzaRating{}, errCreatingPizzaRating
	}

	err = dao.db.QueryRow(
		ctx,
		string(query),
		rating.Restaurant,
		rating.Location,
		rating.DateVisited,
		rating.PizzaOrdered,
		rating.Rating,
		rating.Reccomended,
	).Scan(
		&rating.ID,
		&rating.Restaurant,
		&rating.Location,
		&rating.DateVisited,
		&rating.PizzaOrdered,
		&rating.Rating,
		&rating.Reccomended,
		&rating.CreatedAt,
	)
	if err != nil {
		return models.PizzaRating{}, errCreatingPizzaRating
	}

	return rating, nil
}
