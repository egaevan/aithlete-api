package exercise

import (
	"testing"

	"github.com/aithlete/aithlete-api/domain/entity"
)

func TestListMuscleGroups_Success(t *testing.T) {
	uc := NewListMuscleGroupsUseCase()

	results := uc.ListMuscleGroups()

	if len(results) == 0 {
		t.Fatal("expected non-empty muscle groups")
	}

	foundChest := false
	for _, g := range results {
		if g == string(entity.MuscleGroupChest) {
			foundChest = true
			break
		}
	}
	if !foundChest {
		t.Error("expected chest in muscle groups")
	}
}
