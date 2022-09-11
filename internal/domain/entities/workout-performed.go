package entities

type WorkoutPerformed struct {
	ID                      string `json:"id"`
	UserID                  string `json:"userId"`
	WorkoutID               string `json:"workoutID"`
	Date                    string `json:"date"`
	DayOfWeek               string `json:"dayOfWeek"`
	StartDate               int    `json:"startDate"`
	EndDate                 int    `json:"endDate"`
	TotalTime               int16  `json:"totalTime"`
	Local                   string `json:"local"`
	DescriptionWhenFinished string `json:"descriptionWhenFinished"`
	Exercises               []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Series []struct {
			Executions []struct {
				Load        int `json:"load"`
				Repetitions int `json:"repetitions"`
			} `json:"executions"`
		} `json:"series"`
	} `json:"exercises"`
}
