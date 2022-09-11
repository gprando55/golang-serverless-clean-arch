package repositories

import "application/internal/domain/entities"

type PlanRepository interface {
	Save(plan entities.Plan) error
	FindByUserIdAndName(userId, name string) (plan entities.Plan, err error)
	FindAllByUserId(userId string) (plans []entities.Plan, err error)
}
