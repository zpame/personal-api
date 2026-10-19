INSERT INTO ZackPizzaRatings (
    Restaurant,
    Location,
    DateVisited,
    PizzaOrdered,
    Rating,
    Reccomended
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    id,
    restaurant,
    location,
    datevisited,
    pizzaordered,
    rating,
    reccomended,
    created_at;
