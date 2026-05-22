package dto

type ExerciseDetailResult struct {
	ID           string
	Name         string
	Description  string
	MuscleGroup  string
	Equipment    string
	Difficulty   string
	Instructions []string
	ImageURL     string
	CreatedAt    string
}
