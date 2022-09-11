package entities

type Workout struct {
	ID               string            `json:"id"`
	UserId           string            `json:"userId"`
	ExecutedWorkouts int               `json:"executedWorkouts"`
	Name             string            `json:"name"`
	Exercises        []WorkoutExercise `json:"exercises"`
	CreateDate       int               `json:"createDate"`
	UpdateDate       int               `json:"updateDate"`
}

type WorkoutExercise struct {
	ID                  string          `json:"id"`
	Name                string          `json:"name"`
	NumberOfSeries      int             `json:"numberOfSeries"`
	Series              []WorkoutSeries `json:"series"`
	RepetitionsInterval int             `json:"repetitionsInterval"`
}

type WorkoutSeries struct {
	Index       int `json:"index"`
	Load        int `json:"load"`
	Repetitions int `json:"repetitions"`
}
