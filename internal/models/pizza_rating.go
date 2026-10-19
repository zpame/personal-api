package models

import "time"

type PizzaRating struct {
	ID           int
	Restaurant   string
	Location     string
	DateVisited  time.Time
	PizzaOrdered string
	Rating       float64
	Reccomended  bool
	CreatedAt    time.Time
}
