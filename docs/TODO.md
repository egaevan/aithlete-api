# TODO List - Domain and Application Layer Implementation (DDD + TDD + Clean Architecture)

## Principles

### Clean Architecture Dependency Rule
- Domain layer: pure Go, zero external dependencies, no struct tags
- Application layer: depends only on domain layer (interfaces, not concrete implementations)
- Interfaces layer: depends on application layer (injects use cases via dependency injection)
- Infrastructure layer: implements domain repository interfaces

### DDD (Domain-Driven Design)
- **Entity**: An object with a distinct identity (e.g., User, Workout). Equality is based on ID, not attributes
- **Repository**: Abstraction for data access, defined as interfaces in domain layer
- **Use Case**: Application-specific business rules, orchestrated in application layer
- **DTO**: Data transfer objects for input/output in application layer — no business logic

### TDD (Test-Driven Development) Workflow
For **every** feature, follow the **Red-Green-Refactor** cycle:

1. **Red**: Write a failing test first (before any implementation code)
   └── Domain tests — validate entity behavior and invariants (no mocks)
   └── Use case tests — validate business rules (mock repository interfaces)
2. **Green**: Write the minimum code to make the test pass
3. **Refactor**: Clean up code without changing behavior (tests remain green)

### Test Structure
```text
domain/<entity>/<entity>_test.go    — Unit tests (no external dependencies)
application/usecase/<name>_test.go  — Use case tests (mock repository)
```

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

- [ ] **TDD**: Write failing use case tests for AuthUseCase (mock UserRepository)
- [ ] Define AuthUseCase interface (application ports)
- [ ] Define Auth DTOs (LoginRequest, RegisterRequest, AuthResponse, etc.)
- [ ] Implement AuthUseCase with business logic
  └── Register: hash password, create user, return tokens
  └── Login: verify password, generate JWT tokens
  └── RefreshToken: validate and rotate refresh token
  └── GetMe: fetch user by ID

**Phase 3 — Dependency Injection**

- [ ] Implement UserRepository (infrastructure layer — PostgreSQL)
- [ ] Replace mock provider with real AuthUseCase in AuthHandler

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

- [ ] **TDD**: Write failing use case tests for WorkoutUseCase (mock WorkoutRepository)
- [ ] Define WorkoutUseCase interface
- [ ] Define Workout DTOs (CreateWorkoutRequest, WorkoutResponse, etc.)
- [ ] Implement WorkoutUseCase

**Phase 3 — Dependency Injection**

- [ ] Implement WorkoutRepository (infrastructure layer)
- [ ] Replace mock provider with real WorkoutUseCase in WorkoutHandler

---

### Schedule API

**Phase 1 — Domain Layer**

- [x] Create Schedule domain entity (no tags)
- [x] **TDD**: Write domain tests for Schedule entity
- [x] Define ScheduleRepository interface in domain layer
- [x] Define Schedule domain errors (ErrScheduleNotFound, ErrScheduleConflict)

**Phase 2 — Application Layer**

- [ ] **TDD**: Write failing use case tests for ScheduleUseCase
- [ ] Define ScheduleUseCase interface
- [ ] Define Schedule DTOs
- [ ] Implement ScheduleUseCase

**Phase 3 — Dependency Injection**

- [ ] Implement ScheduleRepository (infrastructure layer)
- [ ] Replace mock provider with real ScheduleUseCase in ScheduleHandler

---

### Goal API

**Phase 1 — Domain Layer**

- [x] Create Goal domain entity (no tags)
- [x] **TDD**: Write domain tests for Goal entity
- [x] Define GoalRepository interface in domain layer
- [x] Define Goal domain errors (ErrGoalNotFound, ErrGoalAlreadyCompleted, ErrInvalidGoalTarget)

**Phase 2 — Application Layer**

- [ ] **TDD**: Write failing use case tests for GoalUseCase
- [ ] Define GoalUseCase interface
- [ ] Define Goal DTOs
- [ ] Implement GoalUseCase

**Phase 3 — Dependency Injection**

- [ ] Implement GoalRepository (infrastructure layer)
- [ ] Replace mock provider with real GoalUseCase in GoalHandler

---

### Progress API

**Phase 1 — Domain Layer**

- [x] Create Progress domain entities (BodyWeight, StrengthProgression, Consistency, etc.)
- [ ] **TDD**: Write domain tests for Progress entities
- [x] Define ProgressRepository interface in domain layer
- [x] Define Progress domain errors (ErrNoData)

**Phase 2 — Application Layer**

- [ ] **TDD**: Write failing use case tests for ProgressUseCase
- [ ] Define ProgressUseCase interface
- [ ] Define Progress DTOs
- [ ] Implement ProgressUseCase

**Phase 3 — Dependency Injection**

- [ ] Implement ProgressRepository (infrastructure layer)
- [ ] Replace mock provider with real ProgressUseCase in ProgressHandler

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

- [x] Profile is part of User entity (domain/user)
- [x] **TDD**: Write domain tests for profile-related behavior

**Phase 2 — Application Layer**

- [ ] **TDD**: Write failing use case tests for ProfileUseCase
- [ ] Define ProfileUseCase interface
- [ ] Define Profile DTOs
- [ ] Implement ProfileUseCase

**Phase 3 — Dependency Injection**

- [ ] Reuse UserRepository for profile operations
- [ ] Replace mock provider with real ProfileUseCase in ProfileHandler

---

## Medium Priority Tasks

### Cross-Cutting Concerns
- [x] Implement domain errors (typed errors for business rule violations)
- [ ] Implement application-level validation
- [ ] Implement database migrations for all entities
- [ ] Add authentication middleware for JWT verification
- [ ] Add request validation middleware
- [ ] Add logging and monitoring for all API endpoints
- [ ] Implement caching for frequently accessed data
- [ ] Add rate limiting for API endpoints
- [ ] Implement API documentation (Swagger/OpenAPI)