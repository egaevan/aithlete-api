package mapper

import (
	"time"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/domain/entity"
)

func ExerciseToResult(e *entity.Exercise) *dto.ExerciseDetailResult {
	return &dto.ExerciseDetailResult{
		ID:           e.ID,
		Name:         e.Name,
		Description:  e.Description,
		MuscleGroup:  e.MuscleGroup,
		Equipment:    e.Equipment,
		Difficulty:   e.Difficulty,
		Instructions: e.Instructions,
		ImageURL:     e.ImageURL,
		CreatedAt:    e.CreatedAt.Format(time.RFC3339),
	}
}

func ExerciseToResultList(exercises []entity.Exercise) []dto.ExerciseDetailResult {
	result := make([]dto.ExerciseDetailResult, len(exercises))
	for i := range exercises {
		result[i] = *ExerciseToResult(&exercises[i])
	}
	return result
}
