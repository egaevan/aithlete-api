package progress

import "github.com/aithlete/aithlete-api/application/usecase"

type Handler struct {
	getBodyWeightHistoryUseCase usecase.GetBodyWeightHistoryUseCase
	addBodyWeightUseCase        usecase.AddBodyWeightUseCase
	getStrengthProgressionUseCase usecase.GetStrengthProgressionUseCase
	getConsistencyUseCase       usecase.GetConsistencyUseCase
	getMuscleVolumeUseCase      usecase.GetMuscleVolumeUseCase
	getProgressOverviewUseCase  usecase.GetProgressOverviewUseCase
}

func New(
	getBodyWeightHistory usecase.GetBodyWeightHistoryUseCase,
	addBodyWeight usecase.AddBodyWeightUseCase,
	getStrengthProgression usecase.GetStrengthProgressionUseCase,
	getConsistency usecase.GetConsistencyUseCase,
	getMuscleVolume usecase.GetMuscleVolumeUseCase,
	getProgressOverview usecase.GetProgressOverviewUseCase,
) *Handler {
	return &Handler{
		getBodyWeightHistoryUseCase:   getBodyWeightHistory,
		addBodyWeightUseCase:          addBodyWeight,
		getStrengthProgressionUseCase: getStrengthProgression,
		getConsistencyUseCase:         getConsistency,
		getMuscleVolumeUseCase:        getMuscleVolume,
		getProgressOverviewUseCase:    getProgressOverview,
	}
}
