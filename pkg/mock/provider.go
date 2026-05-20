package mock

import "time"

func now() string {
	return time.Now().Format(time.RFC3339)
}

type MockProvider struct{}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (p *MockProvider) Login(email, password string) map[string]any {
	return map[string]any{
		"user": map[string]any{
			"id":        "user-001",
			"email":     email,
			"name":      "Alex Johnson",
			"avatar":    "https://api.dicebear.com/7.x/avataaars/svg?seed=Alex",
			"createdAt": "2025-01-01T00:00:00.000Z",
			"updatedAt": "2026-05-20T00:00:00.000Z",
		},
		"tokens": map[string]any{
			"accessToken":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_access_token",
			"refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_refresh_token",
			"expiresIn":    3600,
		},
	}
}

func (p *MockProvider) Register(email, name, password string) map[string]any {
	return map[string]any{
		"user": map[string]any{
			"id":        "user-002",
			"email":     email,
			"name":      name,
			"avatar":    "https://api.dicebear.com/7.x/avataaars/svg?seed=" + name,
			"createdAt": "2026-05-20T00:00:00.000Z",
			"updatedAt": "2026-05-20T00:00:00.000Z",
		},
		"tokens": map[string]any{
			"accessToken":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_access_token",
			"refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_refresh_token",
			"expiresIn":    3600,
		},
	}
}

func (p *MockProvider) Logout() map[string]any {
	return map[string]any{}
}

func (p *MockProvider) GetMe() map[string]any {
	return map[string]any{
		"id":        "user-001",
		"email":     "alex@example.com",
		"name":      "Alex Johnson",
		"avatar":    "https://api.dicebear.com/7.x/avataaars/svg?seed=Alex",
		"createdAt": "2025-01-01T00:00:00.000Z",
		"updatedAt": "2026-05-20T00:00:00.000Z",
		"birthday":  "1995-06-15",
		"gender":    "male",
	}
}

func (p *MockProvider) RefreshToken() map[string]any {
	return map[string]any{
		"accessToken":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_new_access_token",
		"refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_new_refresh_token",
	}
}

func (p *MockProvider) UpdateProfile(name, birthday, gender string) map[string]any {
	return map[string]any{
		"id":        "user-001",
		"email":     "alex@example.com",
		"name":      name,
		"avatar":    "https://api.dicebear.com/7.x/avataaars/svg?seed=" + name,
		"birthday":  birthday,
		"gender":    gender,
		"createdAt": "2025-01-01T00:00:00.000Z",
		"updatedAt": now(),
	}
}

func (p *MockProvider) GetWorkouts() []any {
	return []any{
		map[string]any{
			"id":         "wo-001",
			"name":       "Upper Body Power",
			"date":       "2026-05-19",
			"duration":   52,
			"weightUnit": "lbs",
			"exercises": []any{
				map[string]any{
					"id": "we-001",
					"exercise": map[string]any{
						"id":           "ex-001",
						"name":         "Bench Press",
						"description":  "Barbell bench press for chest",
						"muscleGroup":  "chest",
						"equipment":    "barbell",
						"difficulty":   "intermediate",
						"instructions": []string{"Lie on bench", "Grip barbell", "Lower to chest", "Press up"},
						"createdAt":    "2025-01-01T00:00:00.000Z",
					},
					"sets": []any{
						map[string]any{"id": "s-001", "reps": 10, "weight": 135, "completed": true, "rpe": 7},
						map[string]any{"id": "s-002", "reps": 8, "weight": 155, "completed": true, "rpe": 8},
						map[string]any{"id": "s-003", "reps": 6, "weight": 175, "completed": true, "rpe": 9},
					},
				},
				map[string]any{
					"id": "we-002",
					"exercise": map[string]any{
						"id":           "ex-002",
						"name":         "Overhead Press",
						"description":  "Standing barbell overhead press",
						"muscleGroup":  "shoulders",
						"equipment":    "barbell",
						"difficulty":   "intermediate",
						"instructions": []string{"Stand with feet shoulder-width", "Press bar overhead", "Lower with control"},
						"createdAt":    "2025-01-01T00:00:00.000Z",
					},
					"sets": []any{
						map[string]any{"id": "s-004", "reps": 8, "weight": 95, "completed": true, "rpe": 7},
						map[string]any{"id": "s-005", "reps": 6, "weight": 105, "completed": true, "rpe": 8},
					},
				},
			},
			"notes":        "Felt strong today",
			"completed":    true,
			"createdAt":    "2026-05-19T18:00:00.000Z",
			"updatedAt":    "2026-05-19T19:00:00.000Z",
			"calories":     320,
			"avgHeartRate": 142,
		},
		map[string]any{
			"id":         "wo-002",
			"name":       "Lower Body Strength",
			"date":       "2026-05-17",
			"duration":   65,
			"weightUnit": "lbs",
			"exercises": []any{
				map[string]any{
					"id": "we-003",
					"exercise": map[string]any{
						"id":           "ex-003",
						"name":         "Squat",
						"description":  "Barbell back squat",
						"muscleGroup":  "legs",
						"equipment":    "barbell",
						"difficulty":   "intermediate",
						"instructions": []string{"Position bar on upper back", "Descend below parallel", "Drive through heels"},
						"createdAt":    "2025-01-01T00:00:00.000Z",
					},
					"sets": []any{
						map[string]any{"id": "s-006", "reps": 5, "weight": 225, "completed": true, "rpe": 8},
						map[string]any{"id": "s-007", "reps": 5, "weight": 235, "completed": true, "rpe": 9},
					},
				},
			},
			"notes":        "Heavy day",
			"completed":    true,
			"createdAt":    "2026-05-17T17:00:00.000Z",
			"updatedAt":    "2026-05-17T18:00:00.000Z",
			"calories":     450,
			"avgHeartRate": 155,
		},
	}
}

func (p *MockProvider) GetWorkout(id string) map[string]any {
	return map[string]any{
		"id":         id,
		"name":       "Upper Body Power",
		"date":       "2026-05-19",
		"duration":   52,
		"weightUnit": "lbs",
		"exercises": []any{
			map[string]any{
				"id": "we-001",
				"exercise": map[string]any{
					"id":           "ex-001",
					"name":         "Bench Press",
					"description":  "Barbell bench press for chest",
					"muscleGroup":  "chest",
					"equipment":    "barbell",
					"difficulty":   "intermediate",
					"instructions": []string{"Lie on bench", "Grip barbell", "Lower to chest", "Press up"},
					"createdAt":    "2025-01-01T00:00:00.000Z",
				},
				"sets": []any{
					map[string]any{"id": "s-001", "reps": 10, "weight": 135, "completed": true, "rpe": 7},
					map[string]any{"id": "s-002", "reps": 8, "weight": 155, "completed": true, "rpe": 8},
				},
			},
		},
		"notes":        "Felt strong today",
		"completed":    true,
		"createdAt":    "2026-05-19T18:00:00.000Z",
		"updatedAt":    "2026-05-19T19:00:00.000Z",
		"calories":     320,
		"avgHeartRate": 142,
	}
}

func (p *MockProvider) CreateWorkout(name, date, weightUnit, notes string) map[string]any {
	return map[string]any{
		"id":           "wo-new-001",
		"name":         name,
		"date":         date,
		"duration":     0,
		"weightUnit":   weightUnit,
		"exercises":    []any{},
		"notes":        notes,
		"completed":    false,
		"createdAt":    now(),
		"updatedAt":    now(),
		"calories":     320,
		"avgHeartRate": 142,
	}
}

func (p *MockProvider) UpdateWorkout(id, name, date string) map[string]any {
	return map[string]any{
		"id":         id,
		"name":       name,
		"date":       date,
		"duration":   52,
		"weightUnit": "lbs",
		"exercises":  []any{},
		"completed":  true,
		"createdAt":  "2026-05-19T18:00:00.000Z",
		"updatedAt":  now(),
	}
}

func (p *MockProvider) DeleteWorkout() map[string]any {
	return map[string]any{}
}

func (p *MockProvider) GetWorkoutStats() map[string]any {
	return map[string]any{
		"totalWorkouts": 47,
		"totalVolume":   285000,
		"avgDuration":   54,
		"currentStreak": 14,
	}
}

func (p *MockProvider) GetExercises() []any {
	return []any{
		map[string]any{
			"id":           "ex-001",
			"name":         "Bench Press",
			"description":  "Classic barbell bench press targeting the chest, shoulders, and triceps",
			"muscleGroup":  "chest",
			"equipment":    "barbell",
			"difficulty":   "intermediate",
			"instructions": []string{"Lie flat on bench with eyes under bar", "Grip bar slightly wider than shoulder width", "Unrack and lower bar to mid-chest", "Press bar back up to full arm extension"},
			"imageUrl":     "https://example.com/bench-press.jpg",
			"createdAt":    "2025-01-01T00:00:00.000Z",
		},
		map[string]any{
			"id":           "ex-002",
			"name":         "Overhead Press",
			"description":  "Standing barbell overhead press for shoulder development",
			"muscleGroup":  "shoulders",
			"equipment":    "barbell",
			"difficulty":   "intermediate",
			"instructions": []string{"Stand with feet shoulder-width apart", "Grip bar at shoulder width", "Press bar overhead until arms are fully extended", "Lower with control back to shoulders"},
			"imageUrl":     "https://example.com/ohp.jpg",
			"createdAt":    "2025-01-01T00:00:00.000Z",
		},
		map[string]any{
			"id":           "ex-003",
			"name":         "Squat",
			"description":  "Barbell back squat, the king of lower body exercises",
			"muscleGroup":  "legs",
			"equipment":    "barbell",
			"difficulty":   "intermediate",
			"instructions": []string{"Position bar on upper back", "Stand with feet shoulder-width apart", "Descend by bending knees and hips", "Go below parallel", "Drive through heels to stand"},
			"imageUrl":     "https://example.com/squat.jpg",
			"createdAt":    "2025-01-01T00:00:00.000Z",
		},
		map[string]any{
			"id":           "ex-004",
			"name":         "Romanian Deadlift",
			"description":  "Hip hinge movement targeting hamstrings and glutes",
			"muscleGroup":  "legs",
			"equipment":    "barbell",
			"difficulty":   "intermediate",
			"instructions": []string{"Hold barbell at hip level", "Hinge at hips keeping back straight", "Lower bar along legs", "Feel stretch in hamstrings", "Return to standing"},
			"imageUrl":     "https://example.com/rdl.jpg",
			"createdAt":    "2025-01-01T00:00:00.000Z",
		},
		map[string]any{
			"id":           "ex-005",
			"name":         "Barbell Row",
			"description":  "Bent-over barbell row for back thickness",
			"muscleGroup":  "back",
			"equipment":    "barbell",
			"difficulty":   "intermediate",
			"instructions": []string{"Bend at hips with slight knee bend", "Pull bar to lower chest", "Squeeze shoulder blades", "Lower with control"},
			"imageUrl":     "https://example.com/row.jpg",
			"createdAt":    "2025-01-01T00:00:00.000Z",
		},
	}
}

func (p *MockProvider) GetExercise(id string) map[string]any {
	return map[string]any{
		"id":           id,
		"name":         "Bench Press",
		"description":  "Classic barbell bench press targeting the chest, shoulders, and triceps",
		"muscleGroup":  "chest",
		"equipment":    "barbell",
		"difficulty":   "intermediate",
		"instructions": []string{"Lie flat on bench with eyes under bar", "Grip bar slightly wider than shoulder width", "Unrack and lower bar to mid-chest", "Press bar back up to full arm extension"},
		"imageUrl":     "https://example.com/bench-press.jpg",
		"createdAt":    "2025-01-01T00:00:00.000Z",
	}
}

func (p *MockProvider) GetMuscleGroups() []string {
	return []string{"chest", "back", "legs", "shoulders", "arms", "core", "glutes", "calves", "forearms", "traps", "full-body"}
}

func (p *MockProvider) GetBodyWeightHistory() ([]any, *responseMeta) {
	data := []any{
		map[string]any{"id": "bw-001", "date": "2026-04-08", "weight": 180.0, "bodyFatPercentage": 18.5},
		map[string]any{"id": "bw-002", "date": "2026-04-15", "weight": 179.5, "bodyFatPercentage": 18.2},
		map[string]any{"id": "bw-003", "date": "2026-04-22", "weight": 178.8, "bodyFatPercentage": 17.8},
		map[string]any{"id": "bw-004", "date": "2026-04-29", "weight": 178.2, "bodyFatPercentage": 17.5},
		map[string]any{"id": "bw-005", "date": "2026-05-06", "weight": 177.5, "bodyFatPercentage": 17.2},
		map[string]any{"id": "bw-006", "date": "2026-05-13", "weight": 177.0, "bodyFatPercentage": 16.8},
		map[string]any{"id": "bw-007", "date": "2026-05-20", "weight": 176.5, "bodyFatPercentage": 16.5},
	}
	return data, &responseMeta{Total: 7, Page: 1, Limit: 20, TotalPages: 1}
}

func (p *MockProvider) AddBodyWeight(weight, bodyFat float64) map[string]any {
	return map[string]any{
		"id":                  "bw-new-001",
		"date":                now(),
		"weight":              weight,
		"bodyFatPercentage":   bodyFat,
	}
}

func (p *MockProvider) GetStrengthProgression() ([]any, *responseMeta) {
	data := []any{
		map[string]any{"exercise": "Bench Press", "date": "2026-05-19", "oneRepMax": 185, "volume": 4200},
		map[string]any{"exercise": "Squat", "date": "2026-05-17", "oneRepMax": 225, "volume": 5800},
		map[string]any{"exercise": "Deadlift", "date": "2026-05-15", "oneRepMax": 275, "volume": 4100},
		map[string]any{"exercise": "Overhead Press", "date": "2026-05-19", "oneRepMax": 115, "volume": 2000},
	}
	return data, &responseMeta{Total: 4, Page: 1, Limit: 20, TotalPages: 1}
}

func (p *MockProvider) GetConsistency() ([]any, *responseMeta) {
	data := []any{
		map[string]any{"week": "W1", "workoutsCompleted": 4, "workoutsPlanned": 5, "streak": 4},
		map[string]any{"week": "W2", "workoutsCompleted": 5, "workoutsPlanned": 5, "streak": 9},
		map[string]any{"week": "W3", "workoutsCompleted": 3, "workoutsPlanned": 5, "streak": 3},
		map[string]any{"week": "W4", "workoutsCompleted": 5, "workoutsPlanned": 5, "streak": 8},
		map[string]any{"week": "W5", "workoutsCompleted": 4, "workoutsPlanned": 5, "streak": 4},
		map[string]any{"week": "W6", "workoutsCompleted": 5, "workoutsPlanned": 5, "streak": 14},
	}
	return data, &responseMeta{Total: 6, Page: 1, Limit: 20, TotalPages: 1}
}

func (p *MockProvider) GetMuscleVolume() ([]any, *responseMeta) {
	data := []any{
		map[string]any{"muscleGroup": "Chest", "volume": 4200, "sessions": 8, "trend": "up"},
		map[string]any{"muscleGroup": "Back", "volume": 3800, "sessions": 7, "trend": "up"},
		map[string]any{"muscleGroup": "Legs", "volume": 5100, "sessions": 6, "trend": "stable"},
		map[string]any{"muscleGroup": "Shoulders", "volume": 2800, "sessions": 5, "trend": "up"},
		map[string]any{"muscleGroup": "Arms", "volume": 3200, "sessions": 6, "trend": "down"},
		map[string]any{"muscleGroup": "Core", "volume": 1200, "sessions": 10, "trend": "stable"},
	}
	return data, &responseMeta{Total: 6, Page: 1, Limit: 20, TotalPages: 1}
}

func (p *MockProvider) GetProgressOverview() map[string]any {
	return map[string]any{
		"bodyWeight": []any{
			map[string]any{"id": "bw-001", "date": "2026-04-08", "weight": 180.0},
			map[string]any{"id": "bw-002", "date": "2026-04-15", "weight": 179.5},
			map[string]any{"id": "bw-003", "date": "2026-04-22", "weight": 178.8},
			map[string]any{"id": "bw-004", "date": "2026-04-29", "weight": 178.2},
			map[string]any{"id": "bw-005", "date": "2026-05-06", "weight": 177.5},
			map[string]any{"id": "bw-006", "date": "2026-05-13", "weight": 177.0},
			map[string]any{"id": "bw-007", "date": "2026-05-20", "weight": 176.5},
		},
		"strength": []any{
			map[string]any{"exercise": "Bench Press", "date": "2026-05-19", "oneRepMax": 185, "volume": 4200},
			map[string]any{"exercise": "Squat", "date": "2026-05-17", "oneRepMax": 225, "volume": 5800},
			map[string]any{"exercise": "Deadlift", "date": "2026-05-15", "oneRepMax": 275, "volume": 4100},
		},
		"consistency": []any{
			map[string]any{"week": "W1", "workoutsCompleted": 4, "workoutsPlanned": 5, "streak": 4},
			map[string]any{"week": "W2", "workoutsCompleted": 5, "workoutsPlanned": 5, "streak": 9},
			map[string]any{"week": "W3", "workoutsCompleted": 3, "workoutsPlanned": 5, "streak": 3},
			map[string]any{"week": "W4", "workoutsCompleted": 5, "workoutsPlanned": 5, "streak": 8},
			map[string]any{"week": "W5", "workoutsCompleted": 4, "workoutsPlanned": 5, "streak": 4},
			map[string]any{"week": "W6", "workoutsCompleted": 5, "workoutsPlanned": 5, "streak": 14},
		},
		"muscleVolume": []any{
			map[string]any{"muscleGroup": "Chest", "volume": 4200, "sessions": 8, "trend": "up"},
			map[string]any{"muscleGroup": "Back", "volume": 3800, "sessions": 7, "trend": "up"},
			map[string]any{"muscleGroup": "Legs", "volume": 5100, "sessions": 6, "trend": "stable"},
			map[string]any{"muscleGroup": "Shoulders", "volume": 2800, "sessions": 5, "trend": "up"},
			map[string]any{"muscleGroup": "Arms", "volume": 3200, "sessions": 6, "trend": "down"},
			map[string]any{"muscleGroup": "Core", "volume": 1200, "sessions": 10, "trend": "stable"},
		},
	}
}

func (p *MockProvider) GetAIRecommendations() ([]any, *responseMeta) {
	data := []any{
		map[string]any{
			"id":          "ai-rec-001",
			"type":        "workout",
			"title":       "Upper Body Power",
			"description": "Based on your recovery data, I recommend focusing on upper body today. Your legs need another 24-48 hours of recovery. Consider a chest and back superset workout for maximum efficiency.",
			"confidence":  0.92,
			"createdAt":   "2026-05-20T08:00:00.000Z",
		},
		map[string]any{
			"id":          "ai-rec-002",
			"type":        "recovery",
			"title":       "Active Recovery Day",
			"description": "Your fatigue levels suggest a lighter day tomorrow. Consider mobility work or light cardio.",
			"confidence":  0.85,
			"createdAt":   "2026-05-20T08:00:00.000Z",
		},
		map[string]any{
			"id":          "ai-rec-003",
			"type":        "nutrition",
			"title":       "Increase Protein Intake",
			"description": "Based on your training volume, aim for 1.8g protein per kg bodyweight to support recovery.",
			"confidence":  0.78,
			"createdAt":   "2026-05-20T08:00:00.000Z",
		},
	}
	return data, &responseMeta{Total: 3, Page: 1, Limit: 20, TotalPages: 1}
}

func (p *MockProvider) CreateChatSession() map[string]any {
	return map[string]any{
		"sessionId": "session-1234",
	}
}

func (p *MockProvider) GetChatHistory(sessionID string) ([]any, *responseMeta) {
	data := []any{
		map[string]any{
			"id":        "msg-001",
			"role":      "assistant",
			"content":   "Hi Alex! I've analyzed your recent training data. Your upper body strength is improving nicely, but your legs could use more volume. How can I help you today?",
			"timestamp": "2026-05-20T07:00:00.000Z",
		},
	}
	return data, &responseMeta{Total: 1, Page: 1, Limit: 50, TotalPages: 1}
}

func (p *MockProvider) SendChatMessage(sessionID, message string) ([]any, *responseMeta) {
	data := []any{
		map[string]any{
			"id":        "msg-002",
			"role":      "user",
			"content":   message,
			"timestamp": now(),
		},
		map[string]any{
			"id":        "msg-003",
			"role":      "assistant",
			"content":   "Great question! Based on your current training phase, I'd recommend increasing your bench press volume by adding 2 more working sets. Your recovery data shows your chest is at 92% recovery, so you're well-prepared for additional volume. Try 4 sets of 6-8 reps at 80% of your 1RM.",
			"timestamp": now(),
		},
	}
	return data, &responseMeta{Total: 2, Page: 1, Limit: 50, TotalPages: 1}
}

func (p *MockProvider) GetFatigueAnalysis() map[string]any {
	return map[string]any{
		"overall":    35,
		"central":    40,
		"peripheral": 30,
		"status":     "moderate",
		"factors": []any{
			map[string]any{"name": "Sleep Quality", "value": 85, "impact": "positive"},
			map[string]any{"name": "Training Volume", "value": 65, "impact": "negative"},
			map[string]any{"name": "Nutrition", "value": 70, "impact": "positive"},
			map[string]any{"name": "Stress Level", "value": 55, "impact": "negative"},
			map[string]any{"name": "Hydration", "value": 80, "impact": "positive"},
		},
	}
}

func (p *MockProvider) GetRecoveryScore() map[string]any {
	return map[string]any{
		"overall":    78,
		"sleep":      85,
		"nutrition":  72,
		"stress":     68,
		"muscleRecovery": []any{
			map[string]any{"muscleGroup": "Chest", "recovery": 92, "readyForTraining": true},
			map[string]any{"muscleGroup": "Back", "recovery": 78, "readyForTraining": true},
			map[string]any{"muscleGroup": "Legs", "recovery": 45, "readyForTraining": false},
			map[string]any{"muscleGroup": "Arms", "recovery": 88, "readyForTraining": true},
			map[string]any{"muscleGroup": "Shoulders", "recovery": 70, "readyForTraining": true},
			map[string]any{"muscleGroup": "Core", "recovery": 85, "readyForTraining": true},
		},
		"status": "good",
	}
}

func (p *MockProvider) GetPlateauDetection() ([]any, *responseMeta) {
	data := []any{
		map[string]any{
			"detected":      true,
			"exercise":      "Overhead Press",
			"metric":        "oneRepMax",
			"currentTrend":  "stalled",
			"weeksStalled":  3,
			"suggestions": []string{
				"Try a deload week at 60% volume",
				"Switch to dumbbell variations for 2 weeks",
				"Increase rest between sets to 3-4 minutes",
			},
		},
	}
	return data, &responseMeta{Total: 1, Page: 1, Limit: 20, TotalPages: 1}
}

func (p *MockProvider) GetDashboard() map[string]any {
	return map[string]any{
		"stats": map[string]any{
			"caloriesBurned":    2847,
			"caloriesTrend":     "+12%",
			"activeMinutes":     382,
			"activeMinutesTrend": "+8%",
			"goalsCompleted":    1,
			"goalsTotal":        5,
			"goalsTrend":        "20%",
			"avgHeartRate":      72,
			"heartRateTrend":    "Normal",
		},
		"weeklyProgress": []any{
			map[string]any{"day": "Mon", "calories": 420, "duration": 45},
			map[string]any{"day": "Tue", "calories": 580, "duration": 62},
			map[string]any{"day": "Wed", "calories": 350, "duration": 38},
			map[string]any{"day": "Thu", "calories": 620, "duration": 70},
			map[string]any{"day": "Fri", "calories": 480, "duration": 52},
			map[string]any{"day": "Sat", "calories": 720, "duration": 85},
			map[string]any{"day": "Sun", "calories": 280, "duration": 30},
		},
		"streak": map[string]any{
			"current":      14,
			"personalBest": 21,
			"history":      []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true},
		},
		"recentActivity": []any{
			map[string]any{"id": "act-001", "type": "workout", "title": "Upper Body Power completed", "description": "52 minutes, 5 exercises", "timestamp": "2026-05-19T19:00:00.000Z"},
			map[string]any{"id": "act-002", "type": "achievement", "title": "14 Day Streak!", "description": "You've worked out 14 days in a row", "timestamp": "2026-05-19T19:00:00.000Z"},
			map[string]any{"id": "act-003", "type": "goal", "title": "Volume Goal Completed", "description": "50,000 lbs weekly volume achieved", "timestamp": "2026-05-18T08:00:00.000Z"},
		},
		"muscleRecovery": []any{
			map[string]any{"muscleGroup": "Chest", "recovery": 92, "readyForTraining": true},
			map[string]any{"muscleGroup": "Back", "recovery": 78, "readyForTraining": true},
			map[string]any{"muscleGroup": "Legs", "recovery": 45, "readyForTraining": false},
			map[string]any{"muscleGroup": "Shoulders", "recovery": 70, "readyForTraining": true},
			map[string]any{"muscleGroup": "Arms", "recovery": 88, "readyForTraining": true},
		},
		"todaySchedule": []any{
			map[string]any{"id": "sched-1", "time": "06:30", "title": "Morning Run", "duration": "30 min", "type": "cardio", "completed": true},
			map[string]any{"id": "sched-2", "time": "12:00", "title": "Core Workout", "duration": "20 min", "type": "workout", "completed": true},
			map[string]any{"id": "sched-3", "time": "18:00", "title": "Upper Body", "duration": "45 min", "type": "workout", "completed": false},
		},
		"aiRecommendation": map[string]any{
			"id":          "ai-rec-001",
			"type":        "workout",
			"title":       "Upper Body Power",
			"description": "Based on your recovery data, I recommend focusing on upper body today.",
			"confidence":  0.92,
		},
	}
}

func (p *MockProvider) GetWeeklyProgress() ([]any, *responseMeta) {
	data := []any{
		map[string]any{"day": "Mon", "calories": 420, "duration": 45},
		map[string]any{"day": "Tue", "calories": 580, "duration": 62},
		map[string]any{"day": "Wed", "calories": 350, "duration": 38},
		map[string]any{"day": "Thu", "calories": 620, "duration": 70},
		map[string]any{"day": "Fri", "calories": 480, "duration": 52},
		map[string]any{"day": "Sat", "calories": 720, "duration": 85},
		map[string]any{"day": "Sun", "calories": 280, "duration": 30},
	}
	return data, &responseMeta{Total: 7, Page: 1, Limit: 20, TotalPages: 1}
}

func (p *MockProvider) GetStreak() map[string]any {
	return map[string]any{
		"current":      14,
		"personalBest": 21,
		"history":      []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true},
	}
}

func (p *MockProvider) GetSchedules() []any {
	return []any{
		map[string]any{"id": "sched-1", "date": "2026-05-20", "time": "06:30", "title": "Morning Run", "duration": "30 min", "type": "cardio", "completed": true, "notes": "", "createdAt": "2026-05-19T10:00:00Z", "updatedAt": "2026-05-20T06:30:00Z"},
		map[string]any{"id": "sched-2", "date": "2026-05-20", "time": "12:00", "title": "Core Workout", "duration": "20 min", "type": "workout", "completed": true, "notes": "", "createdAt": "2026-05-19T10:00:00Z", "updatedAt": "2026-05-20T12:00:00Z"},
		map[string]any{"id": "sched-3", "date": "2026-05-20", "time": "18:00", "title": "Upper Body", "duration": "45 min", "type": "workout", "completed": false, "notes": "Focus on chest", "createdAt": "2026-05-19T10:00:00Z", "updatedAt": "2026-05-19T10:00:00Z"},
		map[string]any{"id": "sched-4", "date": "2026-05-21", "time": "07:00", "title": "Yoga Session", "duration": "40 min", "type": "stretching", "completed": false, "notes": "", "createdAt": "2026-05-19T10:00:00Z", "updatedAt": "2026-05-19T10:00:00Z"},
		map[string]any{"id": "sched-5", "date": "2026-05-21", "time": "17:00", "title": "Leg Day", "duration": "50 min", "type": "workout", "completed": false, "notes": "", "createdAt": "2026-05-19T10:00:00Z", "updatedAt": "2026-05-19T10:00:00Z"},
		map[string]any{"id": "sched-6", "date": "2026-05-22", "time": "08:00", "title": "Rest Day", "duration": "0 min", "type": "rest", "completed": false, "notes": "Recovery day", "createdAt": "2026-05-19T10:00:00Z", "updatedAt": "2026-05-19T10:00:00Z"},
	}
}

func (p *MockProvider) GetTodaySchedules() []any {
	return []any{
		map[string]any{"id": "sched-1", "date": "2026-05-20", "time": "06:30", "title": "Morning Run", "duration": "30 min", "type": "cardio", "completed": true, "notes": "", "createdAt": "2026-05-19T10:00:00Z", "updatedAt": "2026-05-20T06:30:00Z"},
		map[string]any{"id": "sched-2", "date": "2026-05-20", "time": "12:00", "title": "Core Workout", "duration": "20 min", "type": "workout", "completed": true, "notes": "", "createdAt": "2026-05-19T10:00:00Z", "updatedAt": "2026-05-20T12:00:00Z"},
		map[string]any{"id": "sched-3", "date": "2026-05-20", "time": "18:00", "title": "Upper Body", "duration": "45 min", "type": "workout", "completed": false, "notes": "Focus on chest", "createdAt": "2026-05-19T10:00:00Z", "updatedAt": "2026-05-19T10:00:00Z"},
	}
}

func (p *MockProvider) GetSchedule(id string) map[string]any {
	return map[string]any{
		"id":        id,
		"date":      "2026-05-20",
		"time":      "06:30",
		"title":     "Morning Run",
		"duration":  "30 min",
		"type":      "cardio",
		"completed": true,
		"notes":     "",
		"createdAt": "2026-05-19T10:00:00Z",
		"updatedAt": "2026-05-20T06:30:00Z",
	}
}

func (p *MockProvider) CreateSchedule(date, time, title, duration, typ, notes string) map[string]any {
	return map[string]any{
		"id":        "sched-new",
		"date":      date,
		"time":      time,
		"title":     title,
		"duration":  duration,
		"type":      typ,
		"completed": false,
		"notes":     notes,
		"createdAt": "2026-05-20T10:00:00Z",
		"updatedAt": "2026-05-20T10:00:00Z",
	}
}

func (p *MockProvider) UpdateSchedule(id, date, time, title, duration, typ, notes string) map[string]any {
	return map[string]any{
		"id":        id,
		"date":      date,
		"time":      time,
		"title":     title,
		"duration":  duration,
		"type":      typ,
		"completed": false,
		"notes":     notes,
		"createdAt": "2026-05-19T10:00:00Z",
		"updatedAt": "2026-05-20T10:00:00Z",
	}
}

func (p *MockProvider) DeleteSchedule() map[string]any {
	return map[string]any{}
}

func (p *MockProvider) ToggleSchedule(id string) map[string]any {
	return map[string]any{
		"id":        id,
		"date":      "2026-05-20",
		"time":      "07:00",
		"title":     "Updated Schedule",
		"duration":  "30 min",
		"type":      "workout",
		"completed": true,
		"notes":     "",
		"createdAt": "2026-05-19T10:00:00Z",
		"updatedAt": "2026-05-20T10:00:00Z",
	}
}

func (p *MockProvider) GetGoals() []any {
	return []any{
		map[string]any{"id": "goal-1", "title": "Complete 20 workouts this month", "type": "workouts", "target": 20, "current": 14, "unit": "sessions", "period": "monthly", "deadline": "2026-05-31", "completed": false, "createdAt": "2026-05-01T00:00:00Z", "updatedAt": "2026-05-18T10:00:00Z"},
		map[string]any{"id": "goal-2", "title": "Burn 10,000 calories this month", "type": "calories", "target": 10000, "current": 7200, "unit": "kcal", "period": "monthly", "deadline": "2026-05-31", "completed": false, "createdAt": "2026-05-01T00:00:00Z", "updatedAt": "2026-05-18T10:00:00Z"},
		map[string]any{"id": "goal-3", "title": "Lift 50,000 lbs total volume this week", "type": "volume", "target": 50000, "current": 50000, "unit": "lbs", "period": "weekly", "deadline": "2026-05-25", "completed": true, "createdAt": "2026-05-19T00:00:00Z", "updatedAt": "2026-05-20T08:00:00Z"},
		map[string]any{"id": "goal-4", "title": "Maintain 7-day workout streak", "type": "streak", "target": 7, "current": 5, "unit": "days", "period": "one-time", "deadline": "", "completed": false, "createdAt": "2026-05-16T00:00:00Z", "updatedAt": "2026-05-20T06:30:00Z"},
		map[string]any{"id": "goal-5", "title": "Run 100 miles this month", "type": "custom", "target": 100, "current": 62, "unit": "miles", "period": "monthly", "deadline": "2026-05-31", "completed": false, "createdAt": "2026-05-01T00:00:00Z", "updatedAt": "2026-05-18T10:00:00Z"},
	}
}

func (p *MockProvider) GetGoal(id string) map[string]any {
	return map[string]any{
		"id":        id,
		"title":     "Complete 20 workouts this month",
		"type":      "workouts",
		"target":    20,
		"current":   14,
		"unit":      "sessions",
		"period":    "monthly",
		"deadline":  "2026-05-31",
		"completed": false,
		"createdAt": "2026-05-01T00:00:00Z",
		"updatedAt": "2026-05-18T10:00:00Z",
	}
}

func (p *MockProvider) CreateGoal(title, typ string, target int, unit, period, deadline string) map[string]any {
	return map[string]any{
		"id":        "goal-new",
		"title":     title,
		"type":      typ,
		"target":    target,
		"current":   0,
		"unit":      unit,
		"period":    period,
		"deadline":  deadline,
		"completed": false,
		"createdAt": "2026-05-20T10:00:00Z",
		"updatedAt": "2026-05-20T10:00:00Z",
	}
}

func (p *MockProvider) UpdateGoal(id, title, typ string, target, current int, unit, period, deadline string) map[string]any {
	return map[string]any{
		"id":        id,
		"title":     title,
		"type":      typ,
		"target":    target,
		"current":   current,
		"unit":      unit,
		"period":    period,
		"deadline":  deadline,
		"completed": false,
		"createdAt": "2026-05-01T00:00:00Z",
		"updatedAt": "2026-05-20T10:00:00Z",
	}
}

func (p *MockProvider) DeleteGoal() map[string]any {
	return map[string]any{}
}

func (p *MockProvider) ToggleGoal(id string) map[string]any {
	return map[string]any{
		"id":        id,
		"title":     "Updated Goal",
		"type":      "workouts",
		"target":    20,
		"current":   20,
		"unit":      "sessions",
		"period":    "monthly",
		"deadline":  "2026-05-31",
		"completed": true,
		"createdAt": "2026-05-01T00:00:00Z",
		"updatedAt": "2026-05-20T10:00:00Z",
	}
}

func (p *MockProvider) UpdateGoalProgress(id string, current int) map[string]any {
	return map[string]any{
		"id":        id,
		"title":     "Updated Goal",
		"type":      "workouts",
		"target":    20,
		"current":   current,
		"unit":      "sessions",
		"period":    "monthly",
		"deadline":  "2026-05-31",
		"completed": false,
		"createdAt": "2026-05-01T00:00:00Z",
		"updatedAt": "2026-05-20T10:00:00Z",
	}
}

func (p *MockProvider) GetAnalyticsOverview() map[string]any {
	return map[string]any{
		"totalVolume":            80100,
		"totalVolumeTrend":       "+15%",
		"avgSession":             58,
		"avgSessionTrend":        "+5 min",
		"sessionsPerMonth":       22,
		"sessionsPerMonthTrend":  "+3",
		"goalCompletion":         20,
		"goalCompletionTrend":    "+4%",
	}
}

func (p *MockProvider) GetWeeklyVolume() []any {
	return []any{
		map[string]any{"week": "W1", "volume": 12500},
		map[string]any{"week": "W2", "volume": 14200},
		map[string]any{"week": "W3", "volume": 13800},
		map[string]any{"week": "W4", "volume": 15600},
		map[string]any{"week": "W5", "volume": 16200},
		map[string]any{"week": "W6", "volume": 17800},
	}
}

func (p *MockProvider) GetMuscleVolumeDistribution() []any {
	return []any{
		map[string]any{"muscle": "Chest", "volume": 4200},
		map[string]any{"muscle": "Back", "volume": 3800},
		map[string]any{"muscle": "Legs", "volume": 5100},
		map[string]any{"muscle": "Shoulders", "volume": 2800},
		map[string]any{"muscle": "Arms", "volume": 3200},
	}
}

type responseMeta struct {
	Total      int
	Page       int
	Limit      int
	TotalPages int
}
