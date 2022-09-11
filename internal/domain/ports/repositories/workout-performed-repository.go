package repositories

import "application/internal/domain/entities"

type WorkoutPerformedRepository interface {
	Save(workoutPerformed entities.WorkoutPerformed) error
}
