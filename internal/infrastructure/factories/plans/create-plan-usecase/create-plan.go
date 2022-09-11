package create_plan

import (
	"application/internal/domain/ports/logger"
	createPlan "application/internal/domain/usecases/plans/create"
	planRepo "application/internal/infrastructure/adapters/repositories/plan"
)

func Make(log logger.Logger) createPlan.CreatePlanUseCase {
	dynamoPlanRepository := planRepo.NewPlanRepository()

	return createPlan.New(dynamoPlanRepository, log)
}
