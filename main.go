package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	DB *dynamodb.DynamoDB
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (interface{}, error) {
	// List all dynamo table
	tablesOutput, err := DB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return nil, err
	}
	return tablesOutput.TableNames, nil
}

func main() {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	DB = dynamodb.New(sess)

	lambda.Start(handler)
}
