package progress

import "github.com/aithlete/aithlete-api/application/usecase/progress"

type Handler struct {
	getBodyWeightHistoryUseCase progress.GetBodyWeightHistoryUseCase
	addBodyWeightUseCase        progress.AddBodyWeightUseCase
	getStrengthProgressionUseCase progress.GetStrengthProgressionUseCase
	getConsistencyUseCase       progress.GetConsistencyUseCase
	getMuscleVolumeUseCase      progress.GetMuscleVolumeUseCase
	getProgressOverviewUseCase  progress.GetProgressOverviewUseCase
}

func New(
	getBodyWeightHistory progress.GetBodyWeightHistoryUseCase,
	addBodyWeight progress.AddBodyWeightUseCase,
	getStrengthProgression progress.GetStrengthProgressionUseCase,
	getConsistency progress.GetConsistencyUseCase,
	getMuscleVolume progress.GetMuscleVolumeUseCase,
	getProgressOverview progress.GetProgressOverviewUseCase,
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
