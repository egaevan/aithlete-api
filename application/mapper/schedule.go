package mapper

import (
	"time"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/domain/entity"
)

func ScheduleToResult(s *entity.Schedule) *dto.ScheduleResult {
	return &dto.ScheduleResult{
		ID:        s.ID,
		UserID:    s.UserID,
		Date:      s.Date,
		Time:      s.Time,
		Title:     s.Title,
		Duration:  s.Duration,
		Type:      s.Type,
		Notes:     s.Notes,
		Completed: s.Completed,
		CreatedAt: s.CreatedAt.Format(time.RFC3339),
		UpdatedAt: s.UpdatedAt.Format(time.RFC3339),
	}
}
