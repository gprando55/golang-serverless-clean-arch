package entities

type Plan struct {
	ID               string        `json:"id"`
	UserId           string        `json:"userId"`
	Name             string        `json:"name"`
	StartDate        int           `json:"startDate"`
	EndDate          int           `json:"endDate"`
	ExecutedWorkouts int32         `json:"executedWorkouts"`
	Workouts         []PlanWorkout `json:"workouts"`
	CreateDate       int           `json:"createDate"`
	UpdateDate       int           `json:"updateDate"`
}

type PlanWorkout struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
