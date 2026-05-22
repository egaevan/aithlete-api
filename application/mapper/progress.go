package mapper

import (
	"time"

	"github.com/aithlete/aithlete-api/application/dto"
	"github.com/aithlete/aithlete-api/domain/entity"
)

func BodyWeightToResult(bw *entity.BodyWeight) *dto.BodyWeightResult {
	return &dto.BodyWeightResult{
		ID:                bw.ID,
		UserID:            bw.UserID,
		Date:              bw.Date,
		Weight:            bw.Weight,
		BodyFatPercentage: bw.BodyFatPercentage,
		CreatedAt:         bw.CreatedAt.Format(time.RFC3339),
	}
}

func StrengthToResult(sr *entity.StrengthRecord) *dto.StrengthResult {
	return &dto.StrengthResult{
		ID:        sr.ID,
		UserID:    sr.UserID,
		Exercise:  sr.Exercise,
		Date:      sr.Date,
		OneRepMax: sr.OneRepMax,
		Volume:    sr.Volume,
	}
}

func ConsistencyToResult(c *entity.Consistency) *dto.ConsistencyResult {
	return &dto.ConsistencyResult{
		Week:              c.Week,
		WorkoutsCompleted: c.WorkoutsCompleted,
		WorkoutsPlanned:   c.WorkoutsPlanned,
		Streak:            c.Streak,
	}
}

func MuscleVolumeToResult(mv *entity.MuscleVolume) *dto.MuscleVolumeResult {
	return &dto.MuscleVolumeResult{
		MuscleGroup: mv.MuscleGroup,
		Volume:      mv.Volume,
		Sessions:    mv.Sessions,
		Trend:       mv.Trend,
	}
}

func BodyWeightToResultList(bws []entity.BodyWeight) []dto.BodyWeightResult {
	result := make([]dto.BodyWeightResult, len(bws))
	for i := range bws {
		result[i] = *BodyWeightToResult(&bws[i])
	}
	return result
}

func StrengthToResultList(srs []entity.StrengthRecord) []dto.StrengthResult {
	result := make([]dto.StrengthResult, len(srs))
	for i := range srs {
		result[i] = *StrengthToResult(&srs[i])
	}
	return result
}

func ConsistencyToResultList(cs []entity.Consistency) []dto.ConsistencyResult {
	result := make([]dto.ConsistencyResult, len(cs))
	for i := range cs {
		result[i] = *ConsistencyToResult(&cs[i])
	}
	return result
}

func MuscleVolumeToResultList(mvs []entity.MuscleVolume) []dto.MuscleVolumeResult {
	result := make([]dto.MuscleVolumeResult, len(mvs))
	for i := range mvs {
		result[i] = *MuscleVolumeToResult(&mvs[i])
	}
	return result
}
