package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	docs "github.com/cdlavacudeg/go-goal-planner/docs"
	"github.com/cdlavacudeg/go-goal-planner/internal/handlers"
)

// @title Goal planner api
// @version 1.0
// @description This is a rest api for goal plannig
// @BasePath /api/v1

// @securityDefinitions.apikey apiKeyAuth
// @in header
// @name Authorization
func main() {
	awsSession, err := setupAwsSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(awsSession)
	// Create table Movies
	tableName := "Users"

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Name"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Email"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Name"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Email"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err = svc.CreateTable(input)
	if err != nil {
		fmt.Printf("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table", tableName)

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
		users := v1.Group("/users")
		{
			users.GET("", handlers.GetUsers)
			users.POST("", handlers.CreateUser)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

// Helloworld godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "Hello World!")
}

func setupAwsSession() (*session.Session, error) {
	// region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region:   aws.String("localhost"),
		Endpoint: aws.String("http://localhost:8000"),
	})
	if err != nil {
		return nil, err
	}

	return awsSession, nil
}
