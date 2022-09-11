package repositories

import "application/internal/domain/entities"

type WorkoutRepository interface {
	Save(workout entities.Workout) error
}
