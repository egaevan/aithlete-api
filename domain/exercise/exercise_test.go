package exercise

import "testing"

func TestNewExercise(t *testing.T) {
	e := NewExercise("Bench Press", "Barbell bench press", "chest", "barbell", "intermediate",
		[]string{"Lie on bench", "Press up"})

	if e.Name != "Bench Press" {
		t.Errorf("expected name Bench Press, got %s", e.Name)
	}
	if e.MuscleGroup != "chest" {
		t.Errorf("expected muscleGroup chest, got %s", e.MuscleGroup)
	}
	if len(e.Instructions) != 2 {
		t.Errorf("expected 2 instructions, got %d", len(e.Instructions))
	}
	if e.CreatedAt.IsZero() {
		t.Error("expected createdAt to be set")
	}
}

func TestValidMuscleGroups(t *testing.T) {
	groups := ValidMuscleGroups()
	if len(groups) == 0 {
		t.Fatal("expected non-empty muscle groups")
	}

	foundChest := false
	for _, g := range groups {
		if g == MuscleGroupChest {
			foundChest = true
			break
		}
	}
	if !foundChest {
		t.Error("expected chest to be in valid muscle groups")
	}
}

func TestNewExerciseEmptyInstructions(t *testing.T) {
	e := NewExercise("Squat", "Barbell squat", "legs", "barbell", "intermediate", nil)

	if e.Instructions != nil {
		t.Error("expected instructions to be nil")
	}
}

func TestMuscleGroupConstants(t *testing.T) {
	if string(MuscleGroupChest) != "chest" {
		t.Errorf("expected chest, got %s", MuscleGroupChest)
	}
	if string(MuscleGroupBack) != "back" {
		t.Errorf("expected back, got %s", MuscleGroupBack)
	}
	if string(MuscleGroupFullBody) != "full-body" {
		t.Errorf("expected full-body, got %s", MuscleGroupFullBody)
	}
}
