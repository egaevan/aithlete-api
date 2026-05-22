package exercise

import "github.com/aithlete/aithlete-api/application/dto"

type ExerciseResponse struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	MuscleGroup  string   `json:"muscleGroup"`
	Equipment    string   `json:"equipment"`
	Difficulty   string   `json:"difficulty"`
	Instructions []string `json:"instructions"`
	ImageURL     string   `json:"imageUrl"`
	CreatedAt    string   `json:"createdAt"`
}

func toExerciseResponse(r *dto.ExerciseDetailResult) ExerciseResponse {
	return ExerciseResponse{
		ID:           r.ID,
		Name:         r.Name,
		Description:  r.Description,
		MuscleGroup:  r.MuscleGroup,
		Equipment:    r.Equipment,
		Difficulty:   r.Difficulty,
		Instructions: r.Instructions,
		ImageURL:     r.ImageURL,
		CreatedAt:    r.CreatedAt,
	}
}

func toExerciseResponseList(results []dto.ExerciseDetailResult) []ExerciseResponse {
	resp := make([]ExerciseResponse, len(results))
	for i, r := range results {
		resp[i] = toExerciseResponse(&r)
	}
	return resp
}

type MuscleGroupResponse struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

func toMuscleGroupResponse(name string) MuscleGroupResponse {
	return MuscleGroupResponse{Name: name, Label: name}
}

func toMuscleGroupResponseList(names []string) []MuscleGroupResponse {
	resp := make([]MuscleGroupResponse, len(names))
	for i, n := range names {
		resp[i] = toMuscleGroupResponse(n)
	}
	return resp
}
