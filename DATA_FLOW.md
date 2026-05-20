# Data Flow & Calculations

This document describes how all dashboard and analytics metrics are collected, calculated, and displayed.

---

## 1. Dashboard Stats

### 1.1 Calories Burned
| Aspect | Detail |
|---|---|
| **Source** | `workouts[].calories` |
| **Input** | User enters calories when creating a workout (`NewWorkoutDialog`) |
| **Field** | `calories` (number, required) |
| **Calculation** | `Σ(workout.calories)` for all workouts in the last 7 days |
| **Trend** | Compare current 7-day total vs previous 7-day total: `((current - previous) / previous) * 100` |
| **Endpoint** | `GET /analytics/dashboard` → `data.stats.caloriesBurned` |
| **Display** | `StatsGrid` component on Dashboard |

### 1.2 Active Minutes
| Aspect | Detail |
|---|---|
| **Source** | `workouts[].duration` (minutes) |
| **Input** | User enters duration when creating a workout (`NewWorkoutDialog`) |
| **Field** | `duration` (number, required, in minutes) |
| **Calculation** | `Σ(workout.duration)` for all workouts in the last 7 days |
| **Trend** | Compare current 7-day total vs previous 7-day total: `((current - previous) / previous) * 100` |
| **Endpoint** | `GET /analytics/dashboard` → `data.stats.activeMinutes` |
| **Display** | `StatsGrid` component on Dashboard |

### 1.3 Goals Completed
| Aspect | Detail |
|---|---|
| **Source** | `goals[]` with `completed` flag |
| **Input** | User creates goals via Goals page (`GoalFormDialog`); marks complete by clicking checkbox |
| **Fields** | `title`, `type`, `target`, `current`, `unit`, `period`, `deadline`, `completed` |
| **Calculation** | `count(goals where completed = true)` / `count(all goals)` |
| **Trend** | Completion rate percentage |
| **Endpoint** | `GET /analytics/dashboard` → `data.stats.goalsCompleted`, `data.stats.goalsTotal` |
| **Display** | `StatsGrid` component on Dashboard (shows "X/Y") |

### 1.4 Avg Heart Rate
| Aspect | Detail |
|---|---|
| **Source** | `workouts[].avgHeartRate` |
| **Input** | User enters heart rate when creating a workout (`NewWorkoutDialog`) |
| **Field** | `avgHeartRate` (number, optional, in bpm) |
| **Calculation** | `AVG(workout.avgHeartRate)` for all workouts in the last 7 days (excluding 0/null) |
| **Trend** | Qualitative label: "Normal", "Elevated", "High" based on range |
| **Endpoint** | `GET /analytics/dashboard` → `data.stats.avgHeartRate` |
| **Display** | `StatsGrid` component on Dashboard (shows value + "bpm") |

---

## 2. Analytics Stats

### 2.1 Total Volume
| Aspect | Detail |
|---|---|
| **Source** | `workouts[].exercises[].sets[]` |
| **Input** | User enters reps + weight per set when creating a workout (`NewWorkoutDialog`) |
| **Calculation** | `Σ(reps × weight)` across all sets in the selected period (default: last 6 weeks) |
| **Trend** | Compare current period vs previous period: `((current - previous) / previous) * 100` |
| **Endpoint** | `GET /analytics/overview` → `data.totalVolume` |
| **Display** | Analytics page summary card |

### 2.2 Avg Session
| Aspect | Detail |
|---|---|
| **Source** | `workouts[].duration` (minutes) |
| **Input** | User enters duration when creating a workout (`NewWorkoutDialog`) |
| **Calculation** | `AVG(workout.duration)` across all workouts in the selected period |
| **Trend** | Compare current period avg vs previous period avg in minutes |
| **Endpoint** | `GET /analytics/overview` → `data.avgSession` |
| **Display** | Analytics page summary card |

### 2.3 Sessions/Month
| Aspect | Detail |
|---|---|
| **Source** | `workouts[]` count |
| **Input** | Derived from workout creation count |
| **Calculation** | `count(workouts)` in the current month |
| **Trend** | Compare current month count vs previous month count |
| **Endpoint** | `GET /analytics/overview` → `data.sessionsPerMonth` |
| **Display** | Analytics page summary card |

### 2.4 Goal Completion
| Aspect | Detail |
|---|---|
| **Source** | `goals[]` with `completed` flag |
| **Input** | User creates goals via Goals page; marks complete manually |
| **Calculation** | `(count(completed goals) / count(all goals)) * 100` |
| **Trend** | Percentage point change vs previous period |
| **Endpoint** | `GET /analytics/overview` → `data.goalCompletion` |
| **Display** | Analytics page summary card |

---

## 3. Workout Input Fields

When creating a workout via `NewWorkoutDialog` (`src/components/workout/new-workout-dialog.tsx`):

| Field | Type | Required | Used For |
|---|---|---|---|
| `name` | string | Yes | Display name |
| `date` | string (YYYY-MM-DD) | Yes | All time-based calculations |
| `duration` | number (minutes) | **Yes** | Active Minutes, Avg Session |
| `calories` | number (kcal) | **Yes** | Calories Burned |
| `avgHeartRate` | number (bpm) | No | Avg Heart Rate |
| `weightUnit` | 'lbs' \| 'kg' | Yes | Volume calculation display |
| `exercises[].exerciseId` | string | Yes | Muscle volume distribution |
| `exercises[].sets[].reps` | number | Yes | Volume calculation |
| `exercises[].sets[].weight` | number | Yes | Volume calculation |
| `notes` | string | No | Display only |

---

## 4. Goal Input Fields

When creating a goal via `GoalFormDialog` (`src/components/goal/goal-form-dialog.tsx`):

| Field | Type | Required | Used For |
|---|---|---|---|
| `title` | string | Yes | Display name |
| `type` | 'workouts' \| 'volume' \| 'calories' \| 'streak' \| 'custom' | Yes | Badge color, auto-completion logic |
| `target` | number | Yes | Completion threshold |
| `unit` | string | Yes | Display unit (e.g., "sessions", "lbs", "kcal") |
| `period` | 'weekly' \| 'monthly' \| 'one-time' | Yes | Badge label, reset period |
| `deadline` | string (YYYY-MM-DD) | No | Deadline display |

### Goal Auto-Completion Logic (backend responsibility)

| Goal Type | Completion Condition |
|---|---|
| `workouts` | `count(workouts in period) >= target` |
| `volume` | `Σ(reps × weight in period) >= target` |
| `calories` | `Σ(calories in period) >= target` |
| `streak` | `currentStreak >= target` |
| `custom` | Manual toggle only |

---

## 5. Schedule Input Fields

When creating a schedule via `ScheduleFormDialog` (`src/components/schedule/schedule-form-dialog.tsx`):

| Field | Type | Required | Used For |
|---|---|---|---|
| `title` | string | Yes | Display name |
| `date` | string (YYYY-MM-DD) | Yes | Date filtering, Today's Schedule |
| `time` | string (HH:MM) | Yes | Time display, sorting |
| `duration` | string | Yes | Display only |
| `type` | 'workout' \| 'cardio' \| 'rest' \| 'stretching' \| 'custom' | Yes | Badge color |
| `notes` | string | No | Display only |

---

## 6. Data Flow Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                          USER INPUT                              │
│                                                                  │
│  ┌──────────────────┐  ┌──────────────────┐  ┌────────────────┐  │
│  │  NewWorkoutDialog│  │   GoalFormDialog │  │ScheduleFormDlg │  │
│  │                  │  │                  │  │                │  │
│  │ • name           │  │ • title          │  │ • title        │  │
│  │ • date           │  │ • type           │  │ • date         │  │
│  │ • duration       │  │ • target         │  │ • time         │  │
│  │ • calories       │  │ • unit           │  │ • duration     │  │
│  │ • avgHeartRate   │  │ • period         │  │ • type         │  │
│  │ • weightUnit     │  │ • deadline       │  │ • notes        │  │
│  │ • exercises[]    │  │                  │  │                │  │
│  │   • exerciseId   │  │                  │  │                │  │
│  │   • sets[]       │  │                  │  │                │  │
│  │     • reps       │  │                  │  │                │  │
│  │     • weight     │  │                  │  │                │  │
│  │ • notes          │  │                  │  │                │  │
│  └────────┬─────────┘  └────────┬─────────┘  └───────┬────────┘  │
│           │                     │                     │           │
│           ▼                     ▼                     ▼           │
│  POST /workouts          POST /goals          POST /schedules     │
└───────────┼───────────────────┼─────────────────────┼─────────────┘
            │                   │                     │
            ▼                   ▼                     ▼
┌─────────────────────────────────────────────────────────────────┐
│                       API / MOCK SERVER                          │
│                                                                   │
│  ┌───────────┐  ┌──────────┐  ┌───────────┐  ┌───────────────┐  │
│  │ workouts  │  │  goals   │  │ schedules │  │   exercises   │  │
│  │  table    │  │  table   │  │   table   │  │    table      │  │
│  │           │  │          │  │           │  │               │  │
│  │ • id      │  │ • id     │  │ • id      │  │ • id          │  │
│  │ • name    │  │ • title  │  │ • date    │  │ • name        │  │
│  │ • date    │  │ • type   │  │ • time    │  │ • muscleGroup │  │
│  │ • duration│  │ • target │  │ • title   │  │ • equipment   │  │
│  │ • calories│  │ • current│  │ • type    │  │ • difficulty  │  │
│  │ • avgHR   │  │ • unit   │  │ • duration│  │               │  │
│  │ • weightU │  │ • period │  │ • notes   │  │               │  │
│  │ • exercis │  │ • deadline│ │ • completed│ │               │  │
│  │ • notes   │  │ • completed│ │           │  │               │  │
│  └───────────┘  └──────────┘  └───────────┘  └───────────────┘  │
└──────────┼───────────────────┼───────────────────┼───────────────┘
           │                   │                   │
           ▼                   ▼                   ▼
┌─────────────────────────────────────────────────────────────────┐
│                      ANALYTICS ENDPOINTS                         │
│                                                                   │
│  GET /analytics/dashboard                                         │
│    → stats.caloriesBurned     = Σ(workout.calories last 7 days)  │
│    → stats.activeMinutes      = Σ(workout.duration last 7 days)  │
│    → stats.goalsCompleted     = count(completed goals)           │
│    → stats.goalsTotal         = count(all goals)                 │
│    → stats.avgHeartRate       = AVG(workout.avgHeartRate)        │
│    → weeklyProgress                                            │
│    → streak                                                    │
│    → muscleRecovery                                            │
│    → todaySchedule        = schedules where date = today        │
│                                                                   │
│  GET /analytics/overview                                          │
│    → totalVolume            = Σ(reps × weight all workouts)      │
│    → avgSession             = AVG(workout.duration)              │
│    → sessionsPerMonth       = count(workouts this month)         │
│    → goalCompletion         = (completed/total) × 100            │
│                                                                   │
│  GET /analytics/volume/weekly  → 6-week volume array             │
│  GET /analytics/volume/muscle  → muscle group volume array       │
│  GET /analytics/weekly                                           │
│  GET /analytics/streak                                           │
└─────────────────────────────────────────────────────────────────┘
           │
           ▼
┌─────────────────────────────────────────────────────────────────┐
│                       FRONTEND DISPLAY                           │
│                                                                   │
│  ┌───────────────────┐  ┌──────────────────┐  ┌───────────────┐  │
│  │    Dashboard      │  │    Analytics     │  │    Goals      │  │
│  │                   │  │                  │  │               │  │
│  │ • StatsGrid       │  │ • Total Volume   │  │ • Goal List   │  │
│  │ • ProgressChart   │  │ • Avg Session    │  │ • Progress    │  │
│  │ • WeeklyAnalytics │  │ • Sessions/Month │  │ • Create/Edit │  │
│  │ • StreakCard      │  │ • Goal Completion│  │ • Delete      │  │
│  │ • MuscleRecovery  │  │ • Weekly Chart   │  │ • Toggle      │  │
│  │ • TodaySchedule   │  │ • Muscle Chart   │  │               │  │
│  │ • AI Coach Panel  │  │                  │  │               │  │
│  └───────────────────┘  └──────────────────┘  └───────────────┘  │
│                                                                   │
│  ┌───────────────────┐  ┌──────────────────┐                     │
│  │    Schedule       │  │    Workouts      │                     │
│  │                   │  │                  │                     │
│  │ • Schedule List   │  │ • Workout List   │                     │
│  │ • Create/Edit     │  │ • New Workout    │                     │
│  │ • Delete          │  │ • Exercises+Sets │                     │
│  │ • Toggle Complete │  │ • Duration       │                     │
│  │ • Today's View    │  │ • Calories       │                     │
│  └───────────────────┘  │ • Heart Rate     │                     │
│                         └──────────────────┘                     │
└─────────────────────────────────────────────────────────────────┘
```

---

## 7. Calculation Formulas Summary

| Metric | Formula | Input Source |
|---|---|---|
| **Set Volume** | `reps × weight` | NewWorkoutDialog → sets |
| **Exercise Volume** | `Σ(set volumes for that exercise)` | NewWorkoutDialog → exercises |
| **Workout Volume** | `Σ(exercise volumes for that workout)` | NewWorkoutDialog → exercises |
| **Total Volume** | `Σ(workout volumes in period)` | Computed from workouts |
| **Calories Burned** | `Σ(workout.calories in last 7 days)` | NewWorkoutDialog → calories |
| **Active Minutes** | `Σ(workout.duration in last 7 days)` | NewWorkoutDialog → duration |
| **Avg Session** | `AVG(workout.duration in period)` | NewWorkoutDialog → duration |
| **Sessions/Month** | `COUNT(workouts in current month)` | Derived from workouts |
| **Goals Completed** | `COUNT(goals where completed = true)` | GoalFormDialog → completed |
| **Goal Completion %** | `(completed / total) × 100` | GoalFormDialog → completed |
| **Avg Heart Rate** | `AVG(workout.avgHeartRate where > 0)` | NewWorkoutDialog → avgHeartRate |
| **Trend %** | `((current - previous) / previous) × 100` | Computed from period comparison |

---

## 8. File Map

| File | Purpose |
|---|---|
| `src/components/workout/new-workout-dialog.tsx` | Input: duration, calories, avgHeartRate, exercises |
| `src/components/goal/goal-form-dialog.tsx` | Input: goal creation/editing |
| `src/components/schedule/schedule-form-dialog.tsx` | Input: schedule creation/editing |
| `src/services/workout.service.ts` | Workout CRUD API calls |
| `src/services/goal.service.ts` | Goal CRUD API calls |
| `src/services/schedule.service.ts` | Schedule CRUD API calls |
| `src/services/analytics.service.ts` | Analytics data aggregation endpoints |
| `src/services/analytics.queries.ts` | TanStack Query hooks for analytics |
| `src/services/goal.queries.ts` | TanStack Query hooks for goals |
| `src/services/schedule.queries.ts` | TanStack Query hooks for schedules |
| `src/app/dashboard/page.tsx` | Dashboard page (uses analytics hooks) |
| `src/app/dashboard/analytics/page.tsx` | Analytics page (uses overview hooks) |
| `src/app/dashboard/goals/page.tsx` | Goals management page |
| `src/app/dashboard/schedule/page.tsx` | Schedule management page |
| `mockoon-aithlete-api.json` | All mock API endpoints |
