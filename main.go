package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (interface{}, error) {
	return event, nil
}

func main() {
	c := dynamo
	lambda.Start(handler)
}
