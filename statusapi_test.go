package main

import (
	"testing"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestIsListRequest(t *testing.T) {
	tests := []struct {
		name string
		request events.APIGatewayProxyRequest
		expect bool
	}{
		{
			"non-list request path",
			events.APIGatewayProxyRequest{Path:"/foo"},
			false,
		},
		{
			"canonical list request path",
			events.APIGatewayProxyRequest{Path:"/instances"},
			true,
		},
		{
			"non-canonical list request path",
			events.APIGatewayProxyRequest{Path:"/instances/"},
			true,
		},
		{
			"non-list request with path params",
			events.APIGatewayProxyRequest{Path:"/instances/",PathParameters: map[string]string{"id":"12345"}},
			false,
		},
		{
			"non-list request with query string params",
			events.APIGatewayProxyRequest{Path:"/instances/",QueryStringParameters: map[string]string{"id":"12345"}},
			false,
		},

	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := isListRequest(test.request)
			assert.Equal(t, test.expect, response)
		})

	}


}