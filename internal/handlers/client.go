package handlers

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	awsConfig "github.com/cdlavacudeg/go-goal-planner/config/aws"
)

var service *dynamodb.DynamoDB

func GetService() *dynamodb.DynamoDB {
	if service != nil {
		return service
	}
	awsSession, err := awsConfig.SetupAwsSession()
	if err != nil {
		panic(err)
	}

	// Set the timeout for the HTTP client
	httpClient := &http.Client{
		Timeout:   time.Second * 1, // Set the desired timeout duration
		Transport: http.DefaultTransport,
	}

	// Create a new AWS config with the configured HTTP client
	awsConfigDynamo := &aws.Config{
		HTTPClient: httpClient,
	}

	// Use the httpClient in your code
	// ...
	service = dynamodb.New(awsSession, awsConfigDynamo)
	// Set service with timeouts

	return service
}
