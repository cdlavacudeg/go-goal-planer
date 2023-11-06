package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/cdlavacudeg/go-goal-planner/utils/env"
)

func SetupAwsSession() (*session.Session, error) {
	// region := os.Getenv("AWS_REGION")
	envConfig, err := env.GetConfig()
	if err != nil {
		return nil, err
	}

	awsSession, err := session.NewSession(&aws.Config{
		Region:   aws.String(envConfig.AwsConfig.Region),
		Endpoint: aws.String(envConfig.DynConfig.Endpoint),
	})
	if err != nil {
		return nil, err
	}

	return awsSession, nil
}
