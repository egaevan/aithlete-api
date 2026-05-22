package workout

import "github.com/aithlete/aithlete-api/application/dto"

type WorkoutResponse struct {
	ID           string                   `json:"id"`
	UserID       string                   `json:"userId"`
	Name         string                   `json:"name"`
	Date         string                   `json:"date"`
	Duration     int                      `json:"duration"`
	WeightUnit   string                   `json:"weightUnit"`
	Notes        string                   `json:"notes"`
	Completed    bool                     `json:"completed"`
	Calories     int                      `json:"calories"`
	AvgHeartRate int                      `json:"avgHeartRate"`
	Exercises    []WorkoutExerciseResponse `json:"exercises"`
	TotalVolume  float64                  `json:"totalVolume"`
	CreatedAt    string                   `json:"createdAt"`
	UpdatedAt    string                   `json:"updatedAt"`
}

type WorkoutExerciseResponse struct {
	ID       string           `json:"id"`
	Exercise ExerciseResponse `json:"exercise"`
	Sets     []SetResponse    `json:"sets"`
}

type ExerciseResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MuscleGroup string `json:"muscleGroup"`
}

type SetResponse struct {
	ID        string  `json:"id"`
	Reps      int     `json:"reps"`
	Weight    float64 `json:"weight"`
	Completed bool    `json:"completed"`
	RPE       int     `json:"rpe"`
}

func toWorkoutResponse(r *dto.WorkoutResult) WorkoutResponse {
	exercises := make([]WorkoutExerciseResponse, len(r.Exercises))
	for i, e := range r.Exercises {
		sets := make([]SetResponse, len(e.Sets))
		for j, s := range e.Sets {
			sets[j] = SetResponse{
				ID:        s.ID,
				Reps:      s.Reps,
				Weight:    s.Weight,
				Completed: s.Completed,
				RPE:       s.RPE,
			}
		}
		exercises[i] = WorkoutExerciseResponse{
			ID: e.ID,
			Exercise: ExerciseResponse{
				ID:          e.Exercise.ID,
				Name:        e.Exercise.Name,
				Description: e.Exercise.Description,
				MuscleGroup: e.Exercise.MuscleGroup,
			},
			Sets: sets,
		}
	}

	return WorkoutResponse{
		ID:           r.ID,
		UserID:       r.UserID,
		Name:         r.Name,
		Date:         r.Date,
		Duration:     r.Duration,
		WeightUnit:   r.WeightUnit,
		Notes:        r.Notes,
		Completed:    r.Completed,
		Calories:     r.Calories,
		AvgHeartRate: r.AvgHeartRate,
		Exercises:    exercises,
		TotalVolume:  r.TotalVolume,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}
}
