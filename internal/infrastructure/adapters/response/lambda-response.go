package response

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type BadRequestInput struct {
	Message    string
	StatusCode int
}

var headers = map[string]string{
	"Access-Control-Allow-Origin": "*",
	"Content-Type":                "application/json",
}

func Ok(data ...interface{}) (events.APIGatewayProxyResponse, error) {

	body, err := json.Marshal(data)

	if err != nil {
		return InternalError()
	}

	bodyString := string(body)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    headers,
		Body:       bodyString,
	}, nil
}

func BadRequest(input BadRequestInput) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: input.StatusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
		Body: "{\"error\":\"" + input.Message + "\"}",
	}, nil

}

func InternalError() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
		Body: "{\"error\":\"Internal server error\"}",
	}, nil

}
