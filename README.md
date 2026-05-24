# Library API

Minimal REST API for personal book library management with soft delete.

## Tech Stack

- Go 1.21+
- Chi router
- PostgreSQL (pgx)
- Docker / Makefile

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/books` | Get all books |
| POST | `/api/books` | Create a book |
| PUT | `/api/books/{id}` | Update a book |
| PATCH | `/api/books/{id}` | Soft delete a book |

## Request/Response Examples

### GET /api/books
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "1984",
    "author": "George Orwell",
    "manufacture": 1949,
    "description": "Dystopian novel"
  }
]
```

### POST /api/books
```json
{
  "title": "Brave New World",
  "author": "Aldous Huxley",
  "manufacture": 1932,
  "description": "Future society novel"
}
```

### Quick Start with Docker
```bash
- make service-up
```

Server runs on http://localhost:8080

## Project Structure

- `cmd/app/` - Entry point
- `internal/`
  - `app/` - Application setup
  - `config/` - Configuration
  - `db/` - Database connection
  - `handlers/` - HTTP handlers
  - `jsonutil/` - JSON utilities
  - `models/` - Data models
  - `repository/` - Database layer
  - `service/` - Business logic
  - `web/` - Frontend (optional)
- `migrations/` - SQL migrations
- `out/` - Build output
