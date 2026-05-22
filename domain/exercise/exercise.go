package exercise

import "time"

type Exercise struct {
	ID           string
	Name         string
	Description  string
	MuscleGroup  string
	Equipment    string
	Difficulty   string
	Instructions []string
	ImageURL     string
	CreatedAt    time.Time
}

type MuscleGroup string

const (
	MuscleGroupChest     MuscleGroup = "chest"
	MuscleGroupBack      MuscleGroup = "back"
	MuscleGroupLegs      MuscleGroup = "legs"
	MuscleGroupShoulders MuscleGroup = "shoulders"
	MuscleGroupArms      MuscleGroup = "arms"
	MuscleGroupCore      MuscleGroup = "core"
	MuscleGroupGlutes    MuscleGroup = "glutes"
	MuscleGroupCalves    MuscleGroup = "calves"
	MuscleGroupForearms  MuscleGroup = "forearms"
	MuscleGroupTraps     MuscleGroup = "traps"
	MuscleGroupFullBody  MuscleGroup = "full-body"
)

func NewExercise(name, description, muscleGroup, equipment, difficulty string, instructions []string) *Exercise {
	return &Exercise{
		Name:         name,
		Description:  description,
		MuscleGroup:  muscleGroup,
		Equipment:    equipment,
		Difficulty:   difficulty,
		Instructions: instructions,
		CreatedAt:    time.Now(),
	}
}

func ValidMuscleGroups() []MuscleGroup {
	return []MuscleGroup{
		MuscleGroupChest, MuscleGroupBack, MuscleGroupLegs,
		MuscleGroupShoulders, MuscleGroupArms, MuscleGroupCore,
		MuscleGroupGlutes, MuscleGroupCalves, MuscleGroupForearms,
		MuscleGroupTraps, MuscleGroupFullBody,
	}
}
