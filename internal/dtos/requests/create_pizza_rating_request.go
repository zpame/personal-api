package requests

type CreatePizzaRatingRequest struct {
	Restaurant   *string  `json:"restaurant"`
	Location     *string  `json:"location"`
	DateVisited  *string  `json:"date_visited"`
	PizzaOrdered *string  `json:"pizza_ordered"`
	Rating       *float64 `json:"rating"`
	Recommend    *bool    `json:"recommend"`
}
