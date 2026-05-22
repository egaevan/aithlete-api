package dto

type CreateWorkoutRequest struct {
	Name       string
	Date       string
	Duration   int
	WeightUnit string
	Notes      string
}

type UpdateWorkoutRequest struct {
	Name string
	Date string
}

type AddExerciseRequest struct {
	ExerciseID   string
	ExerciseName string
	Sets         []SetInput
}

type SetInput struct {
	Reps   int
	Weight float64
	RPE    int
}

type UpdateSetRequest struct {
	Reps   int
	Weight float64
	RPE    int
}

type WorkoutResult struct {
	ID           string
	UserID       string
	Name         string
	Date         string
	Duration     int
	WeightUnit   string
	Notes        string
	Completed    bool
	Calories     int
	AvgHeartRate int
	Exercises    []WorkoutExerciseResult
	TotalVolume  float64
	CreatedAt    string
	UpdatedAt    string
}

type WorkoutExerciseResult struct {
	ID       string
	Exercise ExerciseResult
	Sets     []SetResult
}

type ExerciseResult struct {
	ID          string
	Name        string
	Description string
	MuscleGroup string
}

type SetResult struct {
	ID        string
	Reps      int
	Weight    float64
	Completed bool
	RPE       int
}
