package main

import (
	listPlans "application/internal/domain/usecases/plans/list-by-user"
	"application/internal/infrastructure/adapters/logger"
	"application/internal/infrastructure/adapters/response"
	listPlansUseCaseFactory "application/internal/infrastructure/factories/plans/list-plans-by-user-usecase"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var client = lambda.New(session.New())

func LambdaHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log := logger.NewLogger()

	var params listPlans.Params = listPlans.Params{UserId: request.PathParameters["userId"]}

	log.Info("request to list all plans for user -> ", params)

	useCase := listPlansUseCaseFactory.Make(log)

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
