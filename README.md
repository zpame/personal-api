# personal-api
API for my personal site

## Pizza ratings

`GET /pizza-ratings` returns the pizza ratings.

`POST /pizza-ratings` creates a rating and returns the created record with its
database-assigned ID. Send a JSON object with all of these fields:

```json
{
  "restaurant": "Example Pizzeria",
  "location": "Brooklyn, NY",
  "date_visited": "2026-10-19",
  "pizza_ordered": "Margherita",
  "rating": 8.5,
  "recommend": true
}
```

The restaurant, location, and pizza ordered are required and limited to 100
characters. `date_visited` must use `YYYY-MM-DD`; `rating` must fit the
database's `NUMERIC(4,2)` range. The recommendation value is a required boolean.
