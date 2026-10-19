CREATE TABLE ZackPizzaRatings (
    id SERIAL PRIMARY KEY,
    Restaurant VARCHAR(100) NOT NULL,
    Location VARCHAR(100) NOT NULL,
    DateVisited DATE NOT NULL,
    PizzaOrdered VARCHAR(100) NOT NULL,
    Rating NUMERIC(4,2) NOT NULL,
    Reccomended BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);