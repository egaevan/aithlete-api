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
| Framework | Echo / Fiber |
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

Recommended:
- Clean Architecture
- DDD (Domain Driven Design)
- Hexagonal Architecture

---

# Folder Structure

```text
cmd/

config/

domain/
 ├── user/
 ├── workout/
 ├── exercise/
 ├── ai/
 └── analytics/

application/
 ├── dto/
 ├── usecase/
 ├── mapper/
 └── ports/

interfaces/
 ├── http/
 ├── middleware/
 ├── repository/
 └── router/

infrastructure/
 ├── database/
 ├── ai/
 ├── logger/
 └── cache/

pkg/
```

---

# Modules

# User Module
- Register
- Login
- Refresh token
- User profile
- User goals

Entities:
- User
- Profile

---

# Workout Module
- Create workout
- Add exercise set
- Workout history
- Workout summary
- Weekly volume

Entities:
- Workout
- WorkoutSet
- WorkoutExercise

---

# Exercise Module
- Exercise list
- Muscle group category
- Exercise metadata

Entities:
- Exercise
- ExerciseCategory

---

# AI Module
- Workout recommendation
- Recovery recommendation
- AI chat
- Progressive overload logic

Entities:
- AIRecommendation
- AIConversation

---

# Analytics Module
- Weekly volume analysis
- Plateau detection
- Strength progression
- Recovery scoring

Entities:
- AnalyticsSnapshot

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
