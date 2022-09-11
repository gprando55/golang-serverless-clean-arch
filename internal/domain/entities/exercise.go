package entities

type Exercise struct {
	ID                        string                         `json:"id"`
	UserId                    string                         `json:"userId"`
	Name                      string                         `json:"name"`
	TargetRepetitionsInterval []int                          `json:"targetRepetitionsInterval"`
	MaximumLoadAlreadyDone    ExerciseMaximumLoadAlreadyDone `json:"maximumLoadAlreadyDone"`
	CreateDate                int                            `json:"createDate"`
	UpdateDate                int                            `json:"updateDate"`
}

type ExerciseMaximumLoadAlreadyDone struct {
	Date        int `json:"date"`
	Load        int `json:"load"`
	Repetitions int `json:"repetitions"`
}
