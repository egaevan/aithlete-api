# AITHLETE API

Backend service for the AITHLETE AI fitness application.

## Tech Stack

- **Language**: Golang
- **Framework**: Echo
- **Architecture**: Clean Architecture + DDD
- **Database**: PostgreSQL (future)

## Prerequisites

- Go 1.21+
- Make
- PostgreSQL 15+ (for database features)
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI (for migrations)

```bash
brew install golang-migrate
```

## Setup

### 1. Configuration

Copy `.env.example` to `.env` and adjust values:

```bash
cp .env.example .env
```

### 2. Database

Create the database and run migrations:

```bash
createdb aithlete
make migrate-up
```

### 3. TLS Certificates (optional)

For HTTPS, generate self-signed certificates:

```bash
mkdir -p certs
openssl req -x509 -newkey rsa:4096 -keyout certs/key.pem -out certs/cert.pem \
  -days 365 -nodes -subj "/CN=localhost"
```

Then set `TLS_ENABLED=true` in `.env`.

### 4. Run

```bash
make run
```

The server starts on `http://localhost:8080` (or `https://localhost:8080` if TLS is enabled).

## Development

### Available Commands

```bash
make deps          # Install dependencies (go mod tidy)
make run           # Run the server
make build         # Build the binary
make test          # Run tests
make fmt           # Format code
make lint          # Run golangci-lint
make clean         # Clean build artifacts
make migrate-up    # Run pending database migrations
make migrate-down  # Rollback last migration
make migrate-create# Create a new migration (prompts for name)
```

### Postman Collection

Import `docs/AIthlete API.postman_collection.json` into Postman. All endpoints are pre-configured with:
- `{{baseUrl}}` variable (default: `http://localhost:8080`)
- `{{accessToken}}` auto-populated after Login
- Auth headers on protected endpoints

## Database Migrations

This project uses [golang-migrate/migrate](https://github.com/golang-migrate/migrate) for database migrations.

### Prerequisites

```bash
# Install the migrate CLI
brew install golang-migrate
```

### Commands

```bash
# Run all pending migrations
make migrate-up

# Rollback the last migration
make migrate-down

# Create a new migration
make migrate-create  # prompts for a name
```

Migrations are stored in the `migrations/` directory. The `DATABASE_URL` env var (e.g. `postgres://user:pass@localhost:5432/aithlete?sslmode=disable`) is read from `.env`.

## API Endpoints

All endpoints are prefixed with `/api/v1`.

### Health

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |

### Auth

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/auth/login` | User login |
| POST | `/api/v1/auth/register` | User registration |
| POST | `/api/v1/auth/logout` | User logout |
| GET | `/api/v1/auth/me` | Get current user |
| POST | `/api/v1/auth/refresh` | Refresh token |

### Workouts

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/workouts` | Get workout history |
| POST | `/api/v1/workouts` | Create workout |
| GET | `/api/v1/workouts/:id` | Get single workout |
| PUT | `/api/v1/workouts/:id` | Update workout |
| DELETE | `/api/v1/workouts/:id` | Delete workout |
| GET | `/api/v1/workouts/stats` | Get workout statistics |

### Exercises

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/exercises` | Get exercise list |
| GET | `/api/v1/exercises/:id` | Get single exercise |
| GET | `/api/v1/exercises/muscle-groups` | Get muscle group categories |

### Progress

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/progress/body-weight` | Get body weight history |
| POST | `/api/v1/progress/body-weight` | Add body weight entry |
| GET | `/api/v1/progress/strength` | Get strength progression |
| GET | `/api/v1/progress/consistency` | Get workout consistency |
| GET | `/api/v1/progress/muscle-volume` | Get muscle volume data |
| GET | `/api/v1/progress/overview` | Get full progress overview |

### AI

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/ai/recommendations` | Get AI recommendations |
| POST | `/api/v1/ai/chat` | Create chat session |
| GET | `/api/v1/ai/chat/:sessionId` | Get chat history |
| POST | `/api/v1/ai/chat/:sessionId` | Send chat message |
| GET | `/api/v1/ai/fatigue` | Get fatigue analysis |
| GET | `/api/v1/ai/recovery` | Get recovery score |
| GET | `/api/v1/ai/plateau` | Get plateau detection |

### Analytics

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/analytics/dashboard` | Get dashboard overview |
| GET | `/api/v1/analytics/weekly` | Get weekly progress |
| GET | `/api/v1/analytics/streak` | Get workout streak |
| GET | `/api/v1/analytics/overview` | Get analytics overview |
| GET | `/api/v1/analytics/volume/weekly` | Get weekly training volume |
| GET | `/api/v1/analytics/volume/muscle` | Get muscle volume distribution |

### Schedules

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/schedules` | Get all schedules |
| GET | `/api/v1/schedules/today` | Get today's schedules |
| GET | `/api/v1/schedules/:id` | Get schedule by ID |
| POST | `/api/v1/schedules` | Create schedule |
| PUT | `/api/v1/schedules/:id` | Update schedule |
| DELETE | `/api/v1/schedules/:id` | Delete schedule |
| PATCH | `/api/v1/schedules/:id/toggle` | Toggle completion |

### Goals

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/goals` | Get all goals |
| GET | `/api/v1/goals/:id` | Get goal by ID |
| POST | `/api/v1/goals` | Create goal |
| PUT | `/api/v1/goals/:id` | Update goal |
| DELETE | `/api/v1/goals/:id` | Delete goal |
| PATCH | `/api/v1/goals/:id/toggle` | Toggle completion |
| PATCH | `/api/v1/goals/:id/progress` | Update progress |

### Profile

| Method | Endpoint | Description |
|--------|----------|-------------|
| PUT | `/api/v1/profile` | Update user profile |

## Project Structure

```
cmd/
  server/           # Application entry point
config/             # Configuration files
domain/             # Domain models (DDD)
application/        # Use cases and DTOs
interfaces/
  http/
    handler/        # HTTP handlers
    response/       # Response helpers
    request/        # Request DTOs
  router/           # Route registration
  middleware/       # HTTP middleware
infrastructure/
  config/           # Config loader
  logger/           # Logger
pkg/
  mock/             # Mock data provider
```

## Response Format

All responses follow this structure:

```json
{
  "transaction": {
    "status_code": "00000",
    "status_desc": "Success"
  },
  "data": {},
  "meta": {
    "total": 0,
    "page": 1,
    "limit": 20,
    "total_pages": 1
  }
}
```

## Current Phase: Mock Backend

This is the initial backend integration phase. All endpoints return mock/static data matching the Mockoon API contract. No business logic, database, or AI implementation yet.

### Future Migration

To migrate from mock to real implementation:

1. Replace `pkg/mock/provider.go` with real service implementations
2. Add database connection in `infrastructure/database/`
3. Implement repository layer in `interfaces/repository/`
4. Add use cases in `application/usecase/`
5. Update handlers to call use cases instead of mock provider


