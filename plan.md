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
- **Application** — depends on domain (menggunakan interface repository)
- **Interfaces** — depends on application (injects use case via DI)
- **Infrastructure** — implements domain repository interfaces

### 2. DDD (Domain-Driven Design)
- **Entity**: objek dengan identity unik (User, Workout). Equality based on ID, bukan field values
- **Value Object**: objek immutabel tanpa identitas (Set, Weight)
- **Repository**: interface akses data, didefinisikan di **domain layer**
- **Use Case**: business rules spesifik aplikasi di **application layer**
- **DTO**: data transfer object — hanya untuk input/output, tidak ada business logic

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
domain/<entity>/<entity>_test.go        — Unit tests (pure, no mocks, no deps)
application/usecase/<name>_test.go      — Use case tests (mock repository interfaces)
interfaces/http/handler/<handler>_test.go — Handler tests (httptest, mock use case)
infrastructure/database/<repo>_test.go  — Repository tests (integration, test DB)
```

## Testing Principles
- **Domain layer**: test behavior, business rules, invariants tanpa mock
- **Application layer**: mock repository interfaces, test orchestrasi use case hanya dengan exported/public interfaces
- **No test utilities** seperti testify atau assertion library di domain layer test — cukup `testing.T` dan standard library
- External dependencies hanya di-integrasikan di test infrastructure layer

---

# Folder Structure

```text
cmd/                          — Application entry point

config/                       — Configuration loading (Viper)

domain/                       ── PURE GO, ZERO DEPENDENCIES ──
 ├── user/
 │   ├── user.go              — User entity (no tags)
 │   ├── errors.go            — Domain-specific errors
 │   └── repository.go        — UserRepository interface
 ├── workout/
 │   ├── workout.go           — Workout entity
 │   ├── exercise.go          — WorkoutExercise value object
 │   ├── set.go               — Set value object
 │   ├── errors.go
 │   └── repository.go
 ├── exercise/
 │   ├── exercise.go
 │   ├── errors.go
 │   └── repository.go
 ├── schedule/
 │   ├── schedule.go
 │   ├── errors.go
 │   └── repository.go
 ├── goal/
 │   ├── goal.go
 │   ├── errors.go
 │   └── repository.go
 ├── progress/
 │   ├── bodyweight.go
 │   ├── strength.go
 │   ├── consistency.go
 │   ├── errors.go
 │   └── repository.go
 ├── ai/
 │   ├── recommendation.go
 │   ├── chat.go
 │   ├── fatigue.go
 │   ├── recovery.go
 │   ├── errors.go
 │   └── repository.go
 └── analytics/
     ├── dashboard.go
     ├── volume.go
     ├── errors.go
     └── repository.go

application/                  ── DEPENDS ONLY ON DOMAIN ──
 ├── dto/
 │   ├── auth.go
 │   ├── workout.go
 │   ├── schedule.go
 │   ├── goal.go
 │   ├── progress.go
 │   ├── ai.go
 │   ├── analytics.go
 │   └── exercise.go
 ├── usecase/
 │   ├── auth.go
 │   ├── workout.go
 │   ├── schedule.go
 │   ├── goal.go
 │   ├── progress.go
 │   ├── ai.go
 │   ├── analytics.go
 │   └── exercise.go
 └── ports/
     └── mapper.go

interfaces/                   ── DEPENDS ON APPLICATION ──
 ├── http/
 │   ├── handler/
 │   │   ├── auth.go
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
 │   ├── user_repo.go
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
 │   ├── jwt.go
 │   └── hash.go
 ├── logger/
 │   └── logger.go
 └── cache/
     └── redis.go

pkg/                          — Shared utilities
 └── mock/
     └── provider.go          — Temporary mock, to be removed
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
Domain Entity (domain/)
    │  pure business rules, validates invariants
    ▼
Repository Interface (domain/<entity>/repository.go)
    │  abstraction for data access
    ▼
Repository Implementation (infrastructure/database/)
    │  concrete database operations
```

## Dependency Injection Wire (cmd/)

```go
func main() {
    // Infrastructure
    db := postgres.New(cfg)
    jwtSvc := auth.NewJWT(cfg)
    hashSvc := auth.NewHasher()

    // Repository (infra implements domain interface)
    userRepo := database.NewUserRepository(db)

    // Use Case (application depends on domain interface)
    authUseCase := usecase.NewAuthUseCase(userRepo, jwtSvc, hashSvc)

    // Handler (interfaces depends on use case)
    authHandler := handler.NewAuthHandler(authUseCase)
}
```

---

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
