package lambdaproxy

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ContextKey string

//Executor is an alias for func(context.Context) (interface{}, error). Inside the function, you will find a context with an events.APIGatewayProxyRequest value identified with key "request"
type Executor func(context.Context) (interface{}, error)

//ProxyFunction receives an executor and returns an APIGatewayProxyResponse with json marshalled body and error return
func ProxyFunction(f Executor) func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		statusResponse := http.StatusInternalServerError
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctx = context.WithValue(ctx, ContextKey("request"), &request)

		result, err := f(ctx)
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
