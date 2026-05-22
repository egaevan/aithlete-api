package database

import (
	"context"
	"fmt"
	"time"

	"github.com/aithlete/aithlete-api/domain/entity"
	"github.com/aithlete/aithlete-api/pkg/domainerr"
	"github.com/jackc/pgx/v5"
)

type ExerciseRepository struct {
	pool *Pool
}

func NewExerciseRepository(pool *Pool) *ExerciseRepository {
	return &ExerciseRepository{pool: pool}
}

func (r *ExerciseRepository) FindAll(ctx context.Context) ([]entity.Exercise, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, name, description, muscle_group, equipment, difficulty, instructions, image_url, created_at
		FROM exercises ORDER BY name
	`)
	if err != nil {
		return nil, fmt.Errorf("find all exercises: %w", err)
	}
	defer rows.Close()

	var exercises []entity.Exercise
	for rows.Next() {
		var (
			id, name, description, muscleGroup, equipment, difficulty, imageURL string
			instructions                                                         []string
			createdAt                                                            time.Time
		)
		if err := rows.Scan(&id, &name, &description, &muscleGroup, &equipment, &difficulty, &instructions, &imageURL, &createdAt); err != nil {
			return nil, fmt.Errorf("scan exercise row: %w", err)
		}

		exercises = append(exercises, entity.Exercise{
			ID:           id,
			Name:         name,
			Description:  description,
			MuscleGroup:  muscleGroup,
			Equipment:    equipment,
			Difficulty:   difficulty,
			Instructions: instructions,
			ImageURL:     imageURL,
			CreatedAt:    createdAt,
		})
	}

	return exercises, rows.Err()
}

func (r *ExerciseRepository) FindByID(ctx context.Context, id string) (*entity.Exercise, error) {
	var (
		name, description, muscleGroup, equipment, difficulty, imageURL string
		instructions                                                    []string
		createdAt                                                       time.Time
	)

	err := r.pool.QueryRow(ctx, `
		SELECT id, name, description, muscle_group, equipment, difficulty, instructions, image_url, created_at
		FROM exercises WHERE id = $1
	`, id).Scan(
		&id, &name, &description, &muscleGroup, &equipment, &difficulty, &instructions, &imageURL, &createdAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("find exercise by id %s: %w", id, domainerr.ErrExerciseNotFound)
		}
		return nil, fmt.Errorf("find exercise by id: %w", err)
	}

	return &entity.Exercise{
		ID:           id,
		Name:         name,
		Description:  description,
		MuscleGroup:  muscleGroup,
		Equipment:    equipment,
		Difficulty:   difficulty,
		Instructions: instructions,
		ImageURL:     imageURL,
		CreatedAt:    createdAt,
	}, nil
}

func (r *ExerciseRepository) FindByMuscleGroup(ctx context.Context, muscleGroup string) ([]entity.Exercise, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, name, description, muscle_group, equipment, difficulty, instructions, image_url, created_at
		FROM exercises WHERE muscle_group = $1 ORDER BY name
	`, muscleGroup)
	if err != nil {
		return nil, fmt.Errorf("find exercises by muscle group: %w", err)
	}
	defer rows.Close()

	var exercises []entity.Exercise
	for rows.Next() {
		var (
			id, name, description, mg, equipment, difficulty, imageURL string
			instructions                                                []string
			createdAt                                                   time.Time
		)
		if err := rows.Scan(&id, &name, &description, &mg, &equipment, &difficulty, &instructions, &imageURL, &createdAt); err != nil {
			return nil, fmt.Errorf("scan exercise row: %w", err)
		}

		exercises = append(exercises, entity.Exercise{
			ID:           id,
			Name:         name,
			Description:  description,
			MuscleGroup:  mg,
			Equipment:    equipment,
			Difficulty:   difficulty,
			Instructions: instructions,
			ImageURL:     imageURL,
			CreatedAt:    createdAt,
		})
	}

	return exercises, rows.Err()
}

func (r *ExerciseRepository) Create(ctx context.Context, e *entity.Exercise) error {
	var id string
	err := r.pool.QueryRow(ctx, `
		INSERT INTO exercises (name, description, muscle_group, equipment, difficulty, instructions, image_url, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`, e.Name, e.Description, e.MuscleGroup, e.Equipment, e.Difficulty, e.Instructions, e.ImageURL, e.CreatedAt).Scan(&id)
	if err != nil {
		return fmt.Errorf("create exercise: %w", err)
	}

	e.ID = id
	return nil
}
