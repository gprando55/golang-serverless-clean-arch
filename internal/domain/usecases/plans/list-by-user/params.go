package list_plans_by_user

type Params struct {
	UserId string `json:"userId validate:"required"`
}
