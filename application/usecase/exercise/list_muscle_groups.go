package exercise

import "github.com/aithlete/aithlete-api/domain/entity"

type ListMuscleGroupsUseCase interface {
	ListMuscleGroups() []string
}

type listMuscleGroupsUseCase struct{}

func NewListMuscleGroupsUseCase() ListMuscleGroupsUseCase {
	return &listMuscleGroupsUseCase{}
}

func (uc *listMuscleGroupsUseCase) ListMuscleGroups() []string {
	groups := entity.ValidMuscleGroups()
	result := make([]string, len(groups))
	for i, g := range groups {
		result[i] = string(g)
	}
	return result
}
