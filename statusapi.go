package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"errors"
	"strings"
)

var (
	ErrInvalidInstancePath = errors.New("Invalid path")
	ErrMissingInstanceId = errors.New("Missing parameter")
)

func isListRequest(request events.APIGatewayProxyRequest) bool {
	if len(request.PathParameters) > 0 || len(request.QueryStringParameters)  > 0 {
		return false
	}

	switch request.Path {
	case "/instances":
		return true
	case "/instances/":
		return true
	default:
		return false
	}

}

func getInstanceId(request events.APIGatewayProxyRequest)(string,error) {
	if strings.HasPrefix(request.Path, "/instances") == false {
		return "", ErrInvalidInstancePath
	}

	var paramMap map[string]string = request.QueryStringParameters
	if len(request.PathParameters) > 0 {
		paramMap = request.PathParameters
	}

	id := paramMap["id"]
	switch id {
	case "":
		return "", ErrMissingInstanceId
	default:
		return id, nil
	}
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data %s for request %v.\n", request.RequestContext.RequestID, request)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	// If path is /instances return list
	// If path is /instances/ and length of path params is zero and length of query params is zero return list
	// If path is /instances/<something> verify there's one path param - return the instance identified by that param
	// If path is /instances and there are query params extract id value - return the instance identified by that param



	fmt.Println("request path is", request.Path)

	fmt.Println("path parameters")
	for k,v := range request.PathParameters {
		fmt.Printf("%s -> %s\n",k,v)
	}

	fmt.Println("query parameters")
	for k,v := range request.QueryStringParameters {
		fmt.Printf("%s -> %s\n",k,v)
	}

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}