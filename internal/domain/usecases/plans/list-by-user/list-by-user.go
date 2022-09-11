package list_plans_by_user

import (
	"application/internal/domain/entities"
	"application/internal/domain/errors"
	"application/internal/domain/ports/logger"
	"application/internal/domain/ports/repositories"
)

type ListPlansByUserUseCase interface {
	Execute(params Params) ([]entities.Plan, *errors.AppError)
}

type useCase struct {
	planRepository repositories.PlanRepository
	log            logger.Logger
}

func New(planRepository repositories.PlanRepository, log logger.Logger) ListPlansByUserUseCase {
	return &useCase{planRepository: planRepository, log: log}
}

func (uc *useCase) Execute(params Params) ([]entities.Plan, *errors.AppError) {
	uc.log.Info("find plans for user ", params.UserId)
	plans, err := uc.planRepository.FindAllByUserId(params.UserId)

	if err != nil {
		appError := errors.InternalError(err)
		return nil, appError
	}

	uc.log.Info("number of plans found ", len(plans))

	return plans, nil
}
