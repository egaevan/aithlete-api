package schedule

import "github.com/aithlete/aithlete-api/application/dto"

type ScheduleResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Title     string `json:"title"`
	Duration  string `json:"duration"`
	Type      string `json:"type"`
	Notes     string `json:"notes"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func toScheduleResponse(r *dto.ScheduleResult) ScheduleResponse {
	return ScheduleResponse{
		ID:        r.ID,
		UserID:    r.UserID,
		Date:      r.Date,
		Time:      r.Time,
		Title:     r.Title,
		Duration:  r.Duration,
		Type:      r.Type,
		Notes:     r.Notes,
		Completed: r.Completed,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func toScheduleResponseList(results []dto.ScheduleResult) []ScheduleResponse {
	resp := make([]ScheduleResponse, len(results))
	for i, r := range results {
		resp[i] = toScheduleResponse(&r)
	}
	return resp
}
