package create_plan

import "application/internal/domain/entities"

type Params struct {
	UserId    string                 `json:"userId"`
	Name      string                 `json:"name"`
	StartDate string                 `json:"startDate"`
	EndDate   string                 `json:"endDate"`
	Workouts  []entities.PlanWorkout `json:"workouts"`
}
