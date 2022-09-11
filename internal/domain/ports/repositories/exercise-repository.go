package repositories

import "application/internal/domain/entities"

type ExerciseRepository interface {
	Save(exercise entities.Exercise) error
}
