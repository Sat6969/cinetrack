# 🎬 CineTrack

A REST API for tracking movies, watchlists, and reviews — built with Go, Gin, and GORM, backed by PostgreSQL, and fully containerized with Docker.

CineTrack lets users browse movies, leave reviews, and see live-updated average ratings — with genres, reviews, and users modeled as proper relational associations rather than flat data.

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go |
| Web Framework | [Gin](https://gin-gonic.com/) |
| ORM | [GORM](https://gorm.io/) |
| Database | PostgreSQL |
| Containerization | Docker + Docker Compose |
| Testing | Go `testing` + `httptest` |

---

## Features

- **Movies** — create, read, update, delete, with genre tagging
- **Genres** — many-to-many relationship with movies (a movie can have multiple genres, a genre spans multiple movies)
- **Reviews** — users can review movies; a movie's average rating is recalculated automatically inside a database transaction whenever a new review is added
- **Users** — basic user management with review history
- **Batch movie creation** — insert multiple movies in one request
- **Scoped queries** — filter high-rated or recently released movies
- **Soft delete + hard delete** — movies are soft-deleted by default, with a separate permanent-delete route
- **Stats endpoint** — total movie/user/review counts
- **Environment-based config** — no secrets committed to the repo
- **Graceful shutdown** — in-flight requests finish before the server exits

---

## Project Structure

```
cinetrack/
├── main.go              # entry point, graceful shutdown
├── database/
│   └── database.go      # DB connection, migrations, connection pooling
├── models/
│   └── models.go         # GORM models + scopes
├── handlers/
│   ├── handlers.go       # route handlers
│   └── handlers_test.go  # handler tests
├── middlewares/
│   └── middlewares.go    # auth middleware
├── routes/
│   └── routes.go         # route registration
├── Dockerfile
├── docker-compose.yml
├── .env                  # local secrets (not committed)
└── .gitignore
```

---

## Getting Started

### Option 1 — Run with Docker (recommended)

This spins up both the API and a PostgreSQL database with a single command — no local Go or Postgres installation required.

```bash
git clone https://github.com/Sat6969/cinetrack.git
cd cinetrack
docker-compose up --build
```

The API will be available at `http://localhost:9092`.

### Option 2 — Run locally

Requires Go and a running PostgreSQL instance.

```bash
git clone https://github.com/Sat6969/cinetrack.git
cd cinetrack
go mod download
go run main.go
```

Create a `.env` file in the project root:

```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=moviedb
DB_PORT=5432
```

---

## API Endpoints

### Movies

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/movies` | List all movies |
| `GET` | `/movies/:id` | Get a movie with its genres and reviews |
| `GET` | `/movies/high-rated` | Movies rated above 7 |
| `GET` | `/movies/recent` | Movies released after 2020 |
| `POST` | `/movies` | Create a movie (genres auto-created if new) |
| `POST` | `/movies/batch` | Create multiple movies at once |
| `PUT` | `/movies/:id` | Update a movie |
| `DELETE` | `/movies/:id` | Soft-delete a movie |
| `DELETE` | `/movies/:id/permanent` | Permanently delete a movie |
| `GET` | `/movies/:id/reviews` | Get all reviews for a movie |

### Users

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/users` | List all users |
| `GET` | `/users/:id` | Get a user with their reviews |
| `POST` | `/users` | Create a user |
| `DELETE` | `/users/:id` | Delete a user |

### Reviews

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/reviews` | Create a review (updates the movie's average rating in a transaction) |

### Stats

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/stats` | Total counts of movies, users, and reviews |

Protected routes require an `Authorization` header.

---

## Example Request

**Create a movie:**

```bash
curl -X POST http://localhost:9092/movies \
  -H "Authorization: token" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Inception",
    "rating": 9,
    "release_year": 2010,
    "description": "A mind-bending thriller",
    "genres": ["Sci-Fi", "Thriller"]
  }'
```

**Add a review:**

```bash
curl -X POST http://localhost:9092/reviews \
  -H "Authorization: token" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "movie_id": 1,
    "rating": 9,
    "comment": "Amazing movie!"
  }'
```

---

## Running Tests

```bash
go test ./handlers/
```

---

## Roadmap

- [x] Core REST API (movies, users, reviews, genres)
- [x] PostgreSQL + GORM associations, transactions, scopes
- [x] Dockerized development environment
- [x] Basic handler tests
- [ ] JWT authentication
- [ ] React frontend
- [ ] CI/CD pipeline
- [ ] Cloud deployment

---

## License

MIT
