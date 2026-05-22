package mapper

import (
	"time"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/domain/entity"
)

func WorkoutToResult(w *entity.Workout) *dto.WorkoutResult {
	exercises := make([]dto.WorkoutExerciseResult, len(w.Exercises))
	for i, e := range w.Exercises {
		sets := make([]dto.SetResult, len(e.Sets))
		for j, s := range e.Sets {
			sets[j] = dto.SetResult{
				ID:        s.ID,
				Reps:      s.Reps,
				Weight:    s.Weight,
				Completed: s.Completed,
				RPE:       s.RPE,
			}
		}
		exercises[i] = dto.WorkoutExerciseResult{
			ID: e.ID,
			Exercise: dto.ExerciseResult{
				ID:          e.Exercise.ID,
				Name:        e.Exercise.Name,
				Description: e.Exercise.Description,
				MuscleGroup: e.Exercise.MuscleGroup,
			},
			Sets: sets,
		}
	}

	return &dto.WorkoutResult{
		ID:           w.ID,
		UserID:       w.UserID,
		Name:         w.Name,
		Date:         w.Date,
		Duration:     w.Duration,
		WeightUnit:   w.WeightUnit,
		Notes:        w.Notes,
		Completed:    w.Completed,
		Calories:     w.Calories,
		AvgHeartRate: w.AvgHeartRate,
		Exercises:    exercises,
		TotalVolume:  w.TotalVolume(),
		CreatedAt:    w.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt:    w.GetUpdatedAt().Format(time.RFC3339),
	}
}
