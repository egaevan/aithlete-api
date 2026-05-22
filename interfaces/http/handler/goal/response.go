package goal

import "github.com/aithlete/aithlete-api/application/dto"

type GoalResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Target    int    `json:"target"`
	Current   int    `json:"current"`
	Unit      string `json:"unit"`
	Period    string `json:"period"`
	Deadline  string `json:"deadline"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func toGoalResponse(r *dto.GoalResult) GoalResponse {
	return GoalResponse{
		ID:        r.ID,
		UserID:    r.UserID,
		Title:     r.Title,
		Type:      r.Type,
		Target:    r.Target,
		Current:   r.Current,
		Unit:      r.Unit,
		Period:    r.Period,
		Deadline:  r.Deadline,
		Completed: r.Completed,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func toGoalResponseList(results []dto.GoalResult) []GoalResponse {
	resp := make([]GoalResponse, len(results))
	for i, r := range results {
		resp[i] = toGoalResponse(&r)
	}
	return resp
}
