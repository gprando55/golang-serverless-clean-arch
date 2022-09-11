package create_plan

import (
	"fmt"
	"application/internal/domain/entities"
	"application/internal/domain/errors"
	"application/internal/domain/ports/logger"
	"application/internal/domain/ports/repositories"
	"application/internal/domain/utils/date"
	"application/internal/domain/utils/uuid"
)

type CreatePlanUseCase interface {
	Execute(params Params) (entities.Plan, *errors.AppError)
}

type useCase struct {
	planRepository repositories.PlanRepository
	log            logger.Logger
}

func New(planRepository repositories.PlanRepository, log logger.Logger) CreatePlanUseCase {
	return &useCase{planRepository: planRepository, log: log}
}

func (uc *useCase) Execute(params Params) (entities.Plan, *errors.AppError) {
	_, err := uc.planRepository.FindByUserIdAndName(params.UserId, params.Name)
	if err == nil {
		uc.log.Error("Plan with name for user id already exists", err)
		return entities.Plan{}, errors.BadRequest("Plan with name already exists", err)
	}

	plan := buildEntity(params)

	uc.log.Info("Creating a new plan ", plan)

	err = uc.planRepository.Save(plan)

	if err != nil {
		return entities.Plan{}, errors.InternalError(err)
	}

	return plan, nil
}

func buildEntity(params Params) entities.Plan {
	nowDynamoFormat := date.Now().DynamoFormat()

	fmt.Println("nowDynamoFormat ->", nowDynamoFormat)
	startDate, _ := date.OfPattern(params.StartDate, "02/01/2006")

	endDate, _ := date.OfPattern(params.EndDate, "02/01/2006")

	return entities.Plan{
		ID:               uuid.New(),
		UserId:           params.UserId,
		Name:             params.Name,
		StartDate:        startDate.DynamoFormat(),
		EndDate:          endDate.DynamoFormat(),
		ExecutedWorkouts: 0,
		Workouts:         params.Workouts,
		CreateDate:       nowDynamoFormat,
		UpdateDate:       nowDynamoFormat,
	}
}
