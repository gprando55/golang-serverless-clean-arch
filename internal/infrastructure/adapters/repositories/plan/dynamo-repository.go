package dynamo_plan_repository

import (
	"errors"
	"application/internal/domain/entities"
	"application/internal/domain/ports/repositories"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type dynamoDB struct {
	Api         dynamodbiface.DynamoDBAPI
	Table       string
	userIdIndex string
}

func NewPlanRepository() repositories.PlanRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	return &dynamoDB{
		Api:         svc,
		Table:       "Plan",
		userIdIndex: "userId-index",
	}
}

func (repo *dynamoDB) Save(newPlan entities.Plan) error {

	item, err := dynamodbattribute.MarshalMap(newPlan)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: &repo.Table,
		Item:      item,
	}

	_, err = repo.Api.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}

func (repo *dynamoDB) FindByUserIdAndName(userId string, name string) (plan entities.Plan, err error) {
	keyCond := expression.Key("userId").Equal(expression.Value(userId))
	filter := expression.Name("name").Equal(expression.Value(name))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).WithFilter(filter).Build()

	if err != nil {
		return entities.Plan{}, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(repo.Table),
		IndexName:                 aws.String(repo.userIdIndex),
		KeyConditionExpression:    expr.KeyCondition(),
		Limit:                     aws.Int64(1),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
	}

	out, err := repo.Api.Query(queryInput)

	if err != nil {
		return entities.Plan{}, err
	}

	plans := []entities.Plan{}

	err = dynamodbattribute.UnmarshalListOfMaps(out.Items, &plans)

	if err != nil || len(plans) < 1 {
		return entities.Plan{}, errors.New("No plans found")
	}

	return plans[0], nil
}

func (repo *dynamoDB) FindAllByUserId(userId string) ([]entities.Plan, error) {

	keyCond := expression.Key("userId").Equal(expression.Value(userId))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()

	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(repo.Table),
		IndexName:                 aws.String(repo.userIdIndex),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
	}

	out, err := repo.Api.Query(queryInput)

	if err != nil {
		return nil, err
	}

	var plans []entities.Plan

	err = dynamodbattribute.UnmarshalListOfMaps(out.Items, &plans)

	if err != nil {
		return nil, errors.New("Error on get plans")
	}

	return plans, nil
}
