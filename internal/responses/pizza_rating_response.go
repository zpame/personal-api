package responses

type PizzaRatingResponse struct {
	ID           int     `json:"id"`
	Restaurant   string  `json:"restaurant"`
	Location     string  `json:"location"`
	PizzaOrdered string  `json:"pizza_ordered"`
	Rating       float64 `json:"rating"`
	Reccomend    bool    `json:"recommend"`
}
