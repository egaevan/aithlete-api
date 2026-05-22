# TODO List - Domain and Application Layer Implementation (DDD + TDD + Clean Architecture)

## Principles

### Database Migrations
- Always create `.up.sql` and `.down.sql` migration files when implementing a new API (entity table)
- Migration files go in `migrations/` directory, numbered sequentially (000001, 000002, ...)
- Each migration file has a unique sequence number and a descriptive name (e.g., `000004_create_schedules_table.up.sql`)
- Migrations must include proper indexes, foreign keys with `ON DELETE CASCADE`, and defaults

### Clean Architecture Dependency Rule
- Domain layer: pure Go, zero external dependencies, no struct tags. **Defines all interfaces** (repository, services)
- Application layer: depends only on domain layer (interfaces, not concrete implementations), **no json tags**
- Interfaces layer: depends on application layer (injects use cases via dependency injection), **json tags here**
- Infrastructure layer: implements domain repository and service interfaces

### DDD (Domain-Driven Design)
- **Entity**: An object with a distinct identity (e.g., User, Workout). Equality is based on ID, not attributes
- **Repository**: Abstraction for data access, defined as interfaces in domain layer
- **Use Case**: Application-specific business rules, orchestrated in application layer
- **DTO**: Data transfer objects for input/output in application layer — no business logic, **no json tags**
- **Response Types**: Structs with json tags defined in **handler layer** (mapper from DTO to response)

### TDD (Test-Driven Development) Workflow
For **every** feature, follow the **Red-Green-Refactor** cycle:

1. **Red**: Write a failing test first (before any implementation code)
   └── Domain tests — validate entity behavior and invariants (no mocks)
   └── Use case tests — validate business rules (mock repository interfaces)
2. **Green**: Write the minimum code to make the test pass
3. **Refactor**: Clean up code without changing behavior (tests remain green)

### Test Structure
```text
domain/entity/<entity>_test.go                              — Unit tests (no external dependencies)
domain/repository/<repository>_mock.go                      — Mock repository implementations
domain/service/<service>_mock.go                            — Mock service implementations
application/usecase/<name>_test.go                          — Use case tests (mock repository)
interfaces/http/handler/<domain>/<endpoint>_test.go         — Handler tests (httptest, per endpoint)
```

### Mock Convention
- Mock files use suffix `_mock.go` (not `_test.go`)
- Placed in the **same package** as the interface being mocked
- Exported types so any package can import and use them
- Mocks organized by concept (not by domain module):
  - `domain/repository/user_mock.go` — mock UserRepository
  - `domain/service/token_mock.go` — mock TokenService
  - `domain/service/password_mock.go` — mock PasswordHasher
- Examples:
  - `domain/repository/workout_mock.go` mocks `repository.WorkoutRepository`
  - `domain/service/token_mock.go` mocks `service.TokenService`
  - `domain/service/password_mock.go` mocks `service.PasswordHasher`

---

## API Completion Checklist

- [x] Auth API
- [x] Workout API
- [x] Schedule API
- [x] Progress API
- [x] Profile API
- [x] Goal API
- [ ] AI API
- [ ] Analytics API
- [ ] Exercise API

---

## High Priority Tasks

### Auth API

**Phase 1 — Domain Layer**

- [x] Create User domain entity (no tags, no external dependencies)
- [x] **TDD**: Write domain tests for User entity
  └── TestNewUser validates default values
  └── TestUpdateProfile updates fields
  └── TestUserEquality checks ID-based equality
- [x] Define UserRepository interface (ports) in domain layer
- [x] Define User domain errors (ErrEmailAlreadyExists, ErrInvalidCredentials, ErrUserNotFound)

**Phase 2 — Application Layer**

- [x] **TDD**: Write failing use case tests for AuthUseCase (mock UserRepository)
- [x] Define AuthUseCase interface (application ports)
- [x] Define Auth DTOs (LoginRequest, RegisterRequest, AuthResponse, etc.)
- [x] Implement AuthUseCase with business logic
  └── Register: hash password, create user, return tokens
  └── Login: verify password, generate JWT tokens
  └── RefreshToken: validate and rotate refresh token
  └── GetMe: fetch user by ID

**Phase 3 — Dependency Injection**

- [x] Implement UserRepository (infrastructure layer — PostgreSQL + in-memory dev fallback)
- [x] Replace mock provider with real AuthUseCase in AuthHandler
- [x] Create auth middleware (JWT validation from Authorization header)
- [x] Wire DI in router: config → services → repos → use cases → handlers
- [x] Add `RebuildUser(id, ...)` factory for DDD persistence reconstruction
- [x] Refactor handler layer: per-endpoint files under `handler/auth/`, response types with json tags in handler, no json tags in application/dto
- [x] Move use case interfaces alongside implementations (each `usecase/*.go` exports interface + unexported struct)
- [x] Move infrastructure service interfaces (PasswordHasher, TokenService) to `domain/service/` with mocks alongside
- [x] Move result structs (LoginResult, UserResult, TokenResult) from `ports` to `dto` (the bridge between application & interfaces)
- [x] Remove `application/ports/` package entirely
- [x] Restructure `domain/` folder by DDD concept: `entity/`, `repository/`, `service/` — no more per-module packages
  └── `domain/entity/` — all entities (User, Workout, Exercise, Schedule, Goal, Progress, AI, Analytics)
  └── `domain/repository/` — all repository interfaces + mocks (UserRepository, etc.)
  └── `domain/service/` — all domain service interfaces + mocks (PasswordHasher, TokenService)

---

### Workout API

**Phase 1 — Domain Layer**

- [x] Create Workout domain entity (no tags)
- [x] Create WorkoutExercise domain entity (value object)
- [x] Create Set domain entity (value object)
- [x] **TDD**: Write domain tests for Workout entity
  └── TestCreateWorkout
  └── TestAddExercise
  └── TestUpdateSet
  └── TestCalculateTotalVolume
- [x] Define WorkoutRepository interface in domain layer
- [x] Define Workout domain errors (ErrWorkoutNotFound, ErrEmptyWorkout, etc.)

**Phase 2 — Application Layer**

- [x] **TDD**: Write failing use case tests for WorkoutUseCase (mock WorkoutRepository)
- [x] Define WorkoutUseCase interface
- [x] Define Workout DTOs (CreateWorkoutRequest, WorkoutResponse, etc.)
- [x] Implement WorkoutUseCase

**Phase 3 — Dependency Injection**

- [x] Implement WorkoutRepository (infrastructure layer — PostgreSQL + in-memory dev fallback)
- [x] Replace mock provider with real WorkoutUseCase in WorkoutHandler

---

### Schedule API

**Phase 1 — Domain Layer**

- [x] Create Schedule domain entity (no tags)
- [x] **TDD**: Write domain tests for Schedule entity
- [x] Define ScheduleRepository interface in domain layer
- [x] Define Schedule domain errors (ErrScheduleNotFound, ErrScheduleConflict)

**Phase 2 — Application Layer**

- [x] **TDD**: Write failing use case tests for ScheduleUseCase
- [x] Define ScheduleUseCase interface
- [x] Define Schedule DTOs
- [x] Implement ScheduleUseCase

**Phase 3 — Dependency Injection**

- [x] Implement ScheduleRepository (infrastructure layer)
- [x] Replace mock provider with real ScheduleUseCase in ScheduleHandler

---

### Goal API

**Phase 1 — Domain Layer**

- [x] Create Goal domain entity (no tags)
- [x] **TDD**: Write domain tests for Goal entity
- [x] Define GoalRepository interface in domain layer
- [x] Define Goal domain errors (ErrGoalNotFound, ErrGoalAlreadyCompleted, ErrInvalidGoalTarget)

**Phase 2 — Application Layer**

- [x] **TDD**: Write failing use case tests for GoalUseCase
- [x] Define GoalUseCase interfaces
- [x] Define Goal DTOs
- [x] Implement GoalUseCase

**Phase 3 — Dependency Injection**

- [x] Implement GoalRepository (infrastructure layer — PostgreSQL + mock dev fallback)
- [x] Replace mock provider with real GoalUseCase in GoalHandler

---

### Progress API

**Phase 1 — Domain Layer**

- [x] Create Progress domain entities (BodyWeight, StrengthProgression, Consistency, etc.)
- [x] **TDD**: Write domain tests for Progress entities
- [x] Define ProgressRepository interface in domain layer
- [x] Define Progress domain errors (ErrNoData)

**Phase 2 — Application Layer**

- [x] **TDD**: Write failing use case tests for ProgressUseCase
- [x] Define ProgressUseCase interfaces
- [x] Define Progress DTOs
- [x] Implement ProgressUseCase

**Phase 3 — Dependency Injection**

- [x] Implement ProgressRepository (infrastructure layer — PostgreSQL + mock dev fallback)
- [x] Replace mock provider with real ProgressUseCase in ProgressHandler

---

### AI API

**Phase 1 — Domain Layer**

- [x] Create AI domain entities (Recommendation, ChatSession, Fatigue, RecoveryScore, etc.)
- [ ] **TDD**: Write domain tests for AI entities
- [x] Define AIRepository interface in domain layer
- [x] Define AI domain errors (ErrSessionNotFound)

**Phase 2 — Application Layer**

- [ ] **TDD**: Write failing use case tests for AIUseCase
- [ ] Define AIUseCase interface
- [ ] Define AI DTOs
- [ ] Implement AIUseCase

**Phase 3 — Dependency Injection**

- [ ] Implement AIRepository (infrastructure layer)
- [ ] Replace mock provider with real AIUseCase in AIHandler

---

### Analytics API

**Phase 1 — Domain Layer**

- [x] Create Analytics domain entities (Dashboard, WeeklyVolume, etc.)
- [ ] **TDD**: Write domain tests for Analytics entities
- [x] Define AnalyticsRepository interface in domain layer
- [x] Define Analytics domain errors (ErrNoData)

**Phase 2 — Application Layer**

- [ ] **TDD**: Write failing use case tests for AnalyticsUseCase
- [ ] Define AnalyticsUseCase interface
- [ ] Define Analytics DTOs
- [ ] Implement AnalyticsUseCase

**Phase 3 — Dependency Injection**

- [ ] Implement AnalyticsRepository (infrastructure layer)
- [ ] Replace mock provider with real AnalyticsUseCase in AnalyticsHandler

---

### Exercise API

**Phase 1 — Domain Layer**

- [x] Create Exercise domain entity (no tags)
- [x] **TDD**: Write domain tests for Exercise entity
- [x] Define ExerciseRepository interface in domain layer
- [x] Define Exercise domain errors (ErrExerciseNotFound)

**Phase 2 — Application Layer**

- [ ] **TDD**: Write failing use case tests for ExerciseUseCase
- [ ] Define ExerciseUseCase interface
- [ ] Define Exercise DTOs
- [ ] Implement ExerciseUseCase

**Phase 3 — Dependency Injection**

- [ ] Implement ExerciseRepository (infrastructure layer)
- [ ] Replace mock provider with real ExerciseUseCase in ExerciseHandler

---

### Profile API

**Phase 1 — Domain Layer**

- [x] Profile is part of User entity (domain/entity/user.go)
- [x] **TDD**: Write domain tests for profile-related behavior

**Phase 2 — Application Layer**

- [x] **TDD**: Write failing use case tests for ProfileUseCase
- [x] Define ProfileUseCase interface
- [x] Define Profile DTOs (reuses UserResult from auth)
- [x] Implement ProfileUseCase

**Phase 3 — Dependency Injection**

- [x] Reuse UserRepository for profile operations
- [x] Replace mock provider with real ProfileUseCase in ProfileHandler

---

## Medium Priority Tasks

### Cross-Cutting Concerns
- [x] Implement domain errors (typed errors for business rule violations)
- [x] Implement PasswordHasher (bcrypt) in infrastructure/auth
- [x] Implement TokenService (JWT) in infrastructure/auth
- [x] Implement application-level validation
- [ ] Implement database migrations for all entities
  └── [x] `users` — `migrations/000001_create_users_table.up.sql` (id UUID, email, name, password, avatar, birthday, gender, timestamps)
  └── [x] `workouts` — `migrations/000002_create_workouts_tables.up.sql` (id UUID, user_id FK, name, date, duration, weight_unit, notes, completed, calories, avg_heart_rate, exercises JSONB, timestamps)
  └── [ ] `exercises` — `migrations/000003_create_exercises_table.up.sql` (id UUID, name, description, muscle_group, equipment, difficulty, instructions TEXT[], image_url, timestamps)
  └── [x] `schedules` — `migrations/000004_create_schedules_table.up.sql` (id UUID, user_id FK, date, time, title, duration, type, notes, completed, timestamps)
  └── [x] `goals` — `migrations/000005_create_goals_table.up.sql` (id UUID, user_id FK, title, type, target, unit, period, deadline, current, completed, timestamps)
  └── [x] `progress` — `migrations/000006_create_progress_tables.up.sql` (body_weight: id UUID, user_id FK, weight, body_fat_percentage, date, timestamps; strength_progression: id UUID, user_id FK, exercise_id FK, weight, reps, date, timestamps; consistency: id UUID, user_id FK, date, worked_out BOOLEAN, timestamps; muscle_volume: id UUID, user_id FK, date, muscle_group, total_volume, timestamps)
  └── [ ] `ai` — `migrations/000007_create_ai_tables.up.sql` (chat_sessions: id UUID, user_id FK, title, messages JSONB, timestamps; recommendations: id UUID, user_id FK, type, title, description, priority, expires_at, timestamps)
- [x] Add authentication middleware for JWT verification
- [ ] Add request validation middleware
- [ ] Add logging and monitoring for all API endpoints
- [ ] Implement caching for frequently accessed data
- [ ] Add rate limiting for API endpoints
- [ ] Implement API documentation (Swagger/OpenAPI)