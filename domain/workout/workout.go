package workout

import "time"

type Workout struct {
	ID          string
	UserID      string
	Name        string
	Date        string
	Duration    int
	WeightUnit  string
	Notes       string
	Completed   bool
	Calories    int
	AvgHeartRate int
	Exercises   []WorkoutExercise
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type WorkoutExercise struct {
	ID       string
	Exercise ExerciseRef
	Sets     []Set
}

type ExerciseRef struct {
	ID           string
	Name         string
	Description  string
	MuscleGroup  string
	Equipment    string
	Difficulty   string
	Instructions []string
}

type Set struct {
	ID        string
	Reps      int
	Weight    float64
	Completed bool
	RPE       int
}

func NewWorkout(userID, name, date, weightUnit, notes string) *Workout {
	now := time.Now()
	return &Workout{
		UserID:      userID,
		Name:        name,
		Date:        date,
		WeightUnit:  weightUnit,
		Notes:       notes,
		Completed:   false,
		Exercises:   []WorkoutExercise{},
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (w *Workout) GetCreatedAt() time.Time {
	return w.CreatedAt
}

func (w *Workout) GetUpdatedAt() time.Time {
	return w.UpdatedAt
}

func (w *Workout) AddExercise(exercise WorkoutExercise) error {
	for _, e := range w.Exercises {
		if e.Exercise.ID == exercise.Exercise.ID {
			return ErrDuplicateExercise
		}
	}
	w.Exercises = append(w.Exercises, exercise)
	w.UpdatedAt = time.Now()
	return nil
}

func (w *Workout) UpdateSet(exerciseID, setID string, reps int, weight float64, rpe int) error {
	if reps <= 0 || weight <= 0 {
		return ErrInvalidSetValue
	}
	for i, e := range w.Exercises {
		if e.Exercise.ID == exerciseID {
			for j, s := range e.Sets {
				if s.ID == setID {
					w.Exercises[i].Sets[j].Reps = reps
					w.Exercises[i].Sets[j].Weight = weight
					w.Exercises[i].Sets[j].RPE = rpe
					w.UpdatedAt = time.Now()
					return nil
				}
			}
			return ErrSetNotFound
		}
	}
	return ErrExerciseNotFound
}

func (w *Workout) TotalVolume() float64 {
	var total float64
	for _, e := range w.Exercises {
		for _, s := range e.Sets {
			total += float64(s.Reps) * s.Weight
		}
	}
	return total
}

func (w *Workout) Complete() error {
	if len(w.Exercises) == 0 {
		return ErrEmptyWorkout
	}
	for _, e := range w.Exercises {
		if len(e.Sets) == 0 {
			return ErrIncompleteWorkout
		}
	}
	w.Completed = true
	w.UpdatedAt = time.Now()
	return nil
}

func (w *Workout) Update(name, date string) {
	w.Name = name
	w.Date = date
	w.UpdatedAt = time.Now()
}
