package list_plans_by_user

import (
	"application/internal/domain/ports/logger"
	listPlans "application/internal/domain/usecases/plans/list-by-user"
	planRepo "application/internal/infrastructure/adapters/repositories/plan"
)

func Make(log logger.Logger) listPlans.ListPlansByUserUseCase {
	dynamoPlanRepository := planRepo.NewPlanRepository()

	return listPlans.New(dynamoPlanRepository, log)
}
