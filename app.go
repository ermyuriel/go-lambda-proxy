package lambdaproxy

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type contextKey string

type executor func(*context.Context) (interface{}, error)

func ProxyFunction(f executor) func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		statusResponse := http.StatusBadRequest
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctx = context.WithValue(ctx, contextKey("request"), request)

		result, err := f(&ctx)
		if err == nil {
			statusResponse = http.StatusOK
		}
		encodedResult, encodingErr := json.Marshal(result)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: statusResponse,
				Body:       err.Error(),
			}, encodingErr

		}
		response := events.APIGatewayProxyResponse{
			StatusCode: statusResponse,
			Body:       string(encodedResult),
		}

		return response, err

	}
}
