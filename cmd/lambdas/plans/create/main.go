package main

import (
	"encoding/json"
	createPlan "application/internal/domain/usecases/plans/create"
	"application/internal/infrastructure/adapters/logger"
	"application/internal/infrastructure/adapters/response"
	createPlanUseCaseFactory "application/internal/infrastructure/factories/plans/create-plan-usecase"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var client = lambda.New(session.New())

func LambdaHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log := logger.NewLogger()

	var params createPlan.Params

	json.Unmarshal([]byte(request.Body), &params)

	log.Info("request to create new plan -> ", params)

	useCase := createPlanUseCaseFactory.Make(log)

	plan, appError := useCase.Execute(params)

	if appError != nil {
		log.Info("", appError)
		return response.BadRequest(response.BadRequestInput{Message: appError.Message, StatusCode: int(appError.StatusCode)})
	}

	return response.Ok(plan)
}

func main() {
	runtime.Start(LambdaHandler)
}
