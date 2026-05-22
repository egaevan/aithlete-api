package mapper

import (
	"time"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/domain/entity"
)

func GoalToResult(g *entity.Goal) *dto.GoalResult {
	return &dto.GoalResult{
		ID:        g.ID,
		UserID:    g.UserID,
		Title:     g.Title,
		Type:      g.Type,
		Target:    g.Target,
		Current:   g.Current,
		Unit:      g.Unit,
		Period:    g.Period,
		Deadline:  g.Deadline,
		Completed: g.Completed,
		CreatedAt: g.CreatedAt.Format(time.RFC3339),
		UpdatedAt: g.UpdatedAt.Format(time.RFC3339),
	}
}

func GoalToResultList(goals []entity.Goal) []dto.GoalResult {
	result := make([]dto.GoalResult, len(goals))
	for i := range goals {
		result[i] = *GoalToResult(&goals[i])
	}
	return result
}
