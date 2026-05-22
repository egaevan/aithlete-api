# Gym AI App — Backend Development Plan

# Overview

Backend untuk aplikasi web-based:
- Workout tracking
- AI recommendation
- Progress analytics
- AI fitness assistant

Tech Stack:
- Golang
- PostgreSQL
- OpenCode AI Agent
- Redis/RabbitMQ (future)

---

# Backend Stack

| Component | Technology |
|---|---|
| Language | Golang |
| Framework | Echo |
| Database | PostgreSQL |
| ORM | sqlc / bun |
| Auth | JWT |
| Validation | go-playground/validator |
| AI Communication | HTTP Client |
| Queue | Redis / RabbitMQ |
| Logging | Zerolog |
| Config | Viper |

---

# Architecture

## Three Pillars

### 1. Clean Architecture
Lapisan kode dipisahkan dengan **dependency rule**:
- **Domain** — pure Go, zero dependencies, tidak ada struct tags/serialization logic
- **Application** — depends on domain (menggunakan interface repository), **no json tags**
- **Interfaces** — depends on application (injects use case via DI), **json tags here**
- **Infrastructure** — implements domain repository interfaces

### 2. DDD (Domain-Driven Design)
- **Entity**: objek dengan identity unik (User, Workout). Equality based on ID, bukan field values
- **Value Object**: objek immutabel tanpa identitas (Set, Weight)
- **Repository**: interface akses data, didefinisikan di **domain layer**
- **Use Case**: business rules spesifik aplikasi di **application layer**
- **DTO**: data transfer object — hanya untuk input/output, tidak ada business logic, **no json tags**
- **Response Types**: struct dengan json tags didefinisikan di **handler layer** (bukan application layer)

### 3. Hexagonal Architecture (Ports & Adapters)
- **Ports**: interfaces di domain/application layer
- **Adapters**: implementations di infrastructure layer (PostgreSQL, AI service, etc.)
- Core aplikasi tidak tahu menahu tentang database atau framework HTTP

---

# TDD (Test-Driven Development)

## Workflow — Red/Green/Refactor

Untuk **setiap** fitur, tulis test **sebelum** implementation code:

```
┌────────────────────────────────────────────┐
│              1. RED                         │
│   Write a failing test first                │
│   Domain test: pure unit test, no mock      │
│   Use case test: mock repository interface  │
│   → Test fails (feature not yet implemented)│
└──────────────────────┬─────────────────────┘
                       │
                       ▼
┌────────────────────────────────────────────┐
│              2. GREEN                       │
│   Write minimum code to pass test           │
│   Domain: implement entity logic            │
│   Application: implement use case           │
│   → Test passes                             │
└──────────────────────┬─────────────────────┘
                       │
                       ▼
┌────────────────────────────────────────────┐
│              3. REFACTOR                    │
│   Clean code, extract methods, rename       │
│   No behavior changes — tests still green   │
│   → Test still passes                       │
└────────────────────────────────────────────┘
```

## Test Structure

```text
domain/entity/<entity>_test.go                              — Unit tests (pure, no mocks, no deps)
domain/repository/<repository>_mock.go                      — Mock repository implementations
domain/service/<service>_mock.go                            — Mock service implementations
application/usecase/<name>_test.go                          — Use case tests (mock repository interfaces)
interfaces/http/handler/<domain>/<endpoint>_test.go         — Handler tests (httptest, per endpoint)
infrastructure/database/<repo>_test.go                      — Repository tests (integration, test DB)
```

## Testing Principles
- **Domain layer**: test behavior, business rules, invariants tanpa mock
- **Application layer**: mock repository interfaces, test orchestrasi use case hanya dengan exported/public interfaces
- **Mock files** ditempatkan di package yang sama dengan interface-nya, menggunakan suffix `_mock.go` (bukan `_test.go`) agar bisa diimport oleh test package lain
- **No test utilities** seperti testify atau assertion library di domain layer test — cukup `testing.T` dan standard library
- External dependencies hanya di-integrasikan di test infrastructure layer
- Mocks dipisah berdasarkan konsep: `domain/repository/user_mock.go` untuk mock repository, `domain/service/token_mock.go` untuk mock service

---

# Folder Structure

```text
cmd/                          — Application entry point

config/                       — Configuration loading (Viper)

domain/                       ── PURE GO, ZERO DEPENDENCIES ──
 ├── entity/                  — Domain entities (no tags, no deps)
 │   ├── user.go              — User entity
 │   ├── workout.go           — Workout entity + WorkoutExercise + Set value objects
 │   ├── exercise.go          — Exercise entity + MuscleGroup value object
 │   ├── schedule.go          — Schedule entity
 │   ├── goal.go              — Goal entity
 │   ├── progress.go          — BodyWeight, StrengthRecord entities
 │   ├── ai.go                — Recommendation, ChatSession, FatigueAnalysis, etc.
 │   └── analytics.go         — Dashboard, WeeklyVolume entities
 ├── repository/              — Repository interfaces + mocks
 │   ├── user.go              — UserRepository interface
 │   ├── user_mock.go         — MockUserRepository
 │   ├── workout.go           — WorkoutRepository interface
 │   ├── workout_mock.go      — MockWorkoutRepository
 │   ├── exercise.go          — ExerciseRepository interface
 │   ├── exercise_mock.go     — MockExerciseRepository
 │   ├── schedule.go          — ScheduleRepository interface
 │   ├── schedule_mock.go     — MockScheduleRepository
 │   ├── goal.go              — GoalRepository interface
 │   ├── goal_mock.go         — MockGoalRepository
 │   ├── progress.go          — ProgressRepository interface
 │   ├── progress_mock.go     — MockProgressRepository
 │   ├── ai.go                — AIRepository interface
 │   ├── ai_mock.go           — MockAIRepository
 │   ├── analytics.go         — AnalyticsRepository interface
 │   └── analytics_mock.go    — MockAnalyticsRepository
 └── service/                 — Domain service interfaces + mocks
     ├── token.go             — TokenService interface
     ├── token_mock.go        — MockTokenService
     ├── password.go          — PasswordHasher interface
     └── password_mock.go     — MockPasswordHasher

application/                  ── DEPENDS ONLY ON DOMAIN ──
 ├── dto/                     ── pure data transfer (no json tags). Bridge between application & interfaces ──
 │   ├── auth.go              — LoginResult, UserResult, TokenResult, AuthResponse, etc.
 │   ├── workout.go
 │   ├── schedule.go
 │   ├── goal.go
 │   ├── progress.go
 │   ├── ai.go
 │   ├── analytics.go
 │   └── exercise.go
 ├── usecase/                 ── each file: exported interface + unexported implementation ──
 │   ├── login.go             — LoginUseCase interface + loginUseCase
 │   ├── register.go          — RegisterUseCase interface + registerUseCase
 │   ├── refresh_token.go     — RefreshTokenUseCase interface + refreshTokenUseCase
 │   ├── get_me.go            — GetMeUseCase interface + getMeUseCase
 │   └── ...
 ├── service/
 │   └── auth.go              — GenerateAuthResult, IsNotFound (shared logic)
 └── mapper/
     └── user.go              — UserToResult (domain User → dto.UserResult)

interfaces/                   ── DEPENDS ON APPLICATION ──
 ├── http/
 │   ├── handler/
 │   │   └── auth/            ── per-endpoint files, response types with json tags ──
 │   │       ├── handler.go        — Handler struct + constructor
 │   │       ├── response.go       — LoginResponse, UserResponse, TokenResponse (json tags)
 │   │       ├── login.go          — POST /auth/login
 │   │       ├── register.go       — POST /auth/register
 │   │       ├── logout.go         — POST /auth/logout
 │   │       ├── get_me.go         — GET /auth/me
 │   │       └── refresh_token.go  — POST /auth/refresh
 │   │   ├── workout.go
 │   │   ├── workout.go
 │   │   ├── exercise.go
 │   │   ├── schedule.go
 │   │   ├── goal.go
 │   │   ├── progress.go
 │   │   ├── ai.go
 │   │   ├── analytics.go
 │   │   └── profile.go
 │   ├── request/request.go
 │   └── response/response.go
 ├── middleware/
 │   ├── cors.go
 │   ├── auth.go
 │   └── logger.go
 └── router/router.go

infrastructure/               ── IMPLEMENTS DOMAIN INTERFACES ──
 ├── database/
 │   ├── postgres.go
 │       ├── user_repo.go           — PostgreSQL implementation
    ├── user_repo_memory.go    — In-memory implementation (dev fallback)
 │   ├── workout_repo.go
 │   ├── exercise_repo.go
 │   ├── schedule_repo.go
 │   ├── goal_repo.go
 │   ├── progress_repo.go
 │   ├── ai_repo.go
 │   └── migrations/
 │       └── 001_init.sql
 ├── ai/
 │   └── client.go
 ├── auth/
 │   ├── token.go              — JWT TokenService
 │   └── hash.go               — bcrypt PasswordHasher
 ├── logger/
 │   └── logger.go
 └── cache/
     └── redis.go

pkg/                          — Shared utilities
 ├── app/
 │   └── app.go               — Bootstrap (composition root, wires all deps)
 ├── code/
 │   └── code.go              — Status code constants
 └── mock/
     └── provider.go          — Temporary mock (Workout, Exercise, etc.), to be removed as each module gets wired with DI
```

---

---

# Dependency Flow

```
Controller (interfaces/http/handler)
    │  receives DTO request
    ▼
Use Case (application/usecase)
    │  orchestrates business logic
    ▼
Domain Entity (domain/entity/)
    │  pure business rules, validates invariants
    ▼
Repository Interface (domain/repository/)
    │  abstraction for data access
    ▼
Repository Implementation (infrastructure/database/)
    │  concrete database operations
```

## Composition Root (`pkg/app`)

```go
// pkg/app/app.go — Bootstrap semua dependency
func Bootstrap(log *logger.Logger) Dependencies {
    cfg      := config.Load()
    hashSvc  := auth.NewPasswordHasher()       // infrastructure/auth
    tokenSvc := auth.NewTokenService(cfg.Auth.JWTSecret, cfg.Auth.JWTExpiration)

    pool, err := database.NewPool(ctx, database.DSN(cfg.Database))

    var userRepo repository.UserRepository
    if pool != nil {
        userRepo = database.NewUserRepository(pool)
    } else {
        userRepo = database.NewInMemoryUserRepository()
    }

    return Dependencies{
        Config: cfg,
        Handlers: router.Handlers{
            Auth: authhandler.New(
                usecase.NewLoginUseCase(userRepo, hashSvc, tokenSvc),
                usecase.NewRegisterUseCase(userRepo, hashSvc, tokenSvc),
                usecase.NewRefreshTokenUseCase(tokenSvc),
                usecase.NewGetMeUseCase(userRepo),
            ),
            Workout:   handler.NewWorkoutHandler(provider),  // temporary mock
            ...
        },
    }
}
```

### Clean Architecture Dependency Rule

| Layer | Knows About |
|-------|-------------|
| `cmd/` | everything (composition root) |
| `pkg/app` | infrastructure, application, interfaces |
| `interfaces/router` | only handler types (route registration) |
| `interfaces/http/handler` | application use case interfaces |
| `application/usecase` | domain interfaces (repository + service) |
| `domain/entity` | nothing (pure Go) |
| `domain/repository` | nothing (pure Go) — repository interfaces for data access |
| `domain/service` | nothing (pure Go) — service interfaces for infrastructure concerns |

# Development Sequence (TDD)

**Iterasi per modul**, urutan pengerjaan setiap modul:

1. **Domain Layer (TDD)**:
   - Tulis domain test (RED)
   - Implementasi entity, value object, domain errors (GREEN)
   - Refactor domain code (REFACTOR)
   - Test tetap hijau

2. **Repository Interface (TDD)**:
   - Definisikan interface repository di domain layer
   - Tidak perlu test untuk interface (Go interfaces implicit)

3. **Application Layer (TDD)**:
   - Tulis use case test dengan mock repository (RED)
   - Implementasi use case dengan DTOs (GREEN)
   - Refactor (REFACTOR)
   - Test tetap hijau

4. **Infrastructure Layer**:
   - Implementasi concrete repository (PostgreSQL)
   - Repository integration test (test database)

5. **Interface Layer**:
   - Implementasi handler
   - Handler test (httptest)
   - Inject use case ke handler

6. **Wire DI**:
   - Hubungkan semua layer di cmd/main.go

---

# Modules

## User Module
- Register
- Login
- Refresh token
- User profile

### Domain Layer
- **Entity**: User — aggregate root
- **Value Object**: — (none, User membawa semua data)
- **Repository Interface**: UserRepository (FindByEmail, FindByID, Create, Update)
- **Domain Errors**: ErrEmailAlreadyExists, ErrInvalidCredentials, ErrUserNotFound

### Domain Behavior (test-first)
- Register: validasi email unik, hash password, create user
- Login: verifikasi password, generate token
- Update profile: validasi field opsional (birthday, gender)

---

## Workout Module
- Create workout
- Add exercise set
- Workout history
- Workout summary
- Weekly volume

### Domain Layer
- **Entity**: Workout — aggregate root
- **Value Object**: WorkoutExercise, Set
- **Repository Interface**: WorkoutRepository (FindByID, FindByUserID, Create, Update, Delete)
- **Domain Errors**: ErrWorkoutNotFound, ErrEmptyWorkout, ErrDuplicateExercise

### Domain Behavior (test-first)
- Create workout: default completed=false, update timestamps
- Add exercise: validasi exercise belum ada di workout
- Update set: validasi reps dan weight positive
- Calculate total volume: sum(reps × weight)
- Complete workout: validasi semua set terisi

---

## Exercise Module
- Exercise list
- Muscle group category
- Exercise metadata

### Domain Layer
- **Entity**: Exercise
- **Value Object**: MuscleGroup
- **Repository Interface**: ExerciseRepository (FindAll, FindByID, FindMuscleGroups)
- **Domain Errors**: ErrExerciseNotFound

---

## Schedule Module
- List schedules
- Create schedule
- Toggle complete

### Domain Layer
- **Entity**: Schedule
- **Repository Interface**: ScheduleRepository (FindByUserID, FindByID, Create, Update, Delete)
- **Domain Errors**: ErrScheduleNotFound, ErrScheduleConflict

---

## Goal Module
- List goals
- Create goal
- Track progress

### Domain Layer
- **Entity**: Goal
- **Repository Interface**: GoalRepository (FindByUserID, FindByID, Create, Update, Delete)
- **Domain Errors**: ErrGoalNotFound, ErrGoalAlreadyCompleted

---

## Progress Module
- Body weight history
- Strength progression
- Consistency tracking
- Muscle volume

### Domain Layer
- **Entity**: BodyWeight, StrengthRecord
- **Value Object**: Consistency, MuscleVolume
- **Repository Interface**: ProgressRepository (FindBodyWeightByUserID, FindStrengthByUserID, FindConsistency, FindMuscleVolume)
- **Domain Errors**: ErrNoData

---

## AI Module
- Workout recommendation
- Recovery recommendation
- AI chat
- Progressive overload logic
- Fatigue analysis
- Recovery score
- Plateau detection

### Domain Layer
- **Entity**: Recommendation, ChatSession, FatigueAnalysis, RecoveryScore, PlateauDetection
- **Repository Interface**: AIRepository (GetRecommendations, CreateChatSession, GetChatHistory, SendChatMessage, GetFatigueAnalysis, GetRecoveryScore, GetPlateauDetection)
- **Domain Errors**: ErrSessionNotFound

---

## Analytics Module
- Dashboard overview
- Weekly volume
- Streak tracking
- Muscle volume distribution

### Domain Layer
- **Entity**: Dashboard, WeeklyVolume
- **Repository Interface**: AnalyticsRepository (GetDashboard, GetWeeklyProgress, GetStreak, GetWeeklyVolume, GetMuscleVolumeDistribution)
- **Domain Errors**: ErrNoData

---

# Database Schema

## users

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE,
    password TEXT,
    created_at TIMESTAMP
);
```

---

## profiles

```sql
CREATE TABLE profiles (
    id UUID PRIMARY KEY,
    user_id UUID,
    height NUMERIC,
    weight NUMERIC,
    goal TEXT,
    experience_level TEXT
);
```

---

## workouts

```sql
CREATE TABLE workouts (
    id UUID PRIMARY KEY,
    user_id UUID,
    workout_date DATE,
    notes TEXT,
    created_at TIMESTAMP
);
```

---

## workout_sets

```sql
CREATE TABLE workout_sets (
    id UUID PRIMARY KEY,
    workout_id UUID,
    exercise_id UUID,
    reps INTEGER,
    weight NUMERIC,
    rir INTEGER,
    metadata JSONB
);
```

---

## exercises

```sql
CREATE TABLE exercises (
    id UUID PRIMARY KEY,
    name TEXT,
    muscle_group TEXT,
    equipment TEXT
);
```

---

# AI Flow

```text
Frontend
   ↓
Backend API
   ↓
AI Service
   ↓
OpenCode Agent
```

---

# API Endpoints

## Auth

```text
POST   /api/v1/auth/register
POST   /api/v1/auth/login
POST   /api/v1/auth/refresh
```

---

## Workout

```text
GET    /api/v1/workouts
POST   /api/v1/workouts
GET    /api/v1/workouts/:id
PUT    /api/v1/workouts/:id
DELETE /api/v1/workouts/:id
```

---

## Exercise

```text
GET    /api/v1/exercises
GET    /api/v1/exercises/:id
```

---

## AI

```text
POST   /api/v1/ai/recommendation
POST   /api/v1/ai/chat
```

---

## Analytics

```text
GET    /api/v1/analytics/volume
GET    /api/v1/analytics/progression
GET    /api/v1/analytics/recovery
```

---

# Security

- JWT access token
- Refresh token
- Rate limiting
- Request validation
- SQL injection prevention
- Secure cookie

---

# Deployment

## Frontend
- Vercel

## Backend
- Fly.io
- Railway
- Render

---

# Monitoring

| Purpose | Tool |
|---|---|
| Logs | Grafana Loki |
| Metrics | Prometheus |
| Error Tracking | Sentry |
| Uptime | BetterStack |

---

# Future Improvements

- Nutrition tracking
- Meal recommendation
- Smart wearable integration
- Apple Health integration
- Google Fit integration
- AI generated workout plan PDF
- Social/community features
