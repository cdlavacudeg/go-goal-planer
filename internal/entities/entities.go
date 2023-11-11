package entities

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type User struct {
	PK    string `json:"PK"`
	SK    string `json:"SK"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{
		PK:    "USER",
		SK:    "USER",
		Id:    2,
		Name:  "John",
		Email: "j@j.com",
	},
	{
		PK:    "USER",
		SK:    "USER",
		Id:    2,
		Name:  "Jane",
		Email: "jane@j.com",
	},
}

func GetUsers() []User {
	return users
}

func GetUser(id int) (*User, error) {
	for _, user := range users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func CreateUser(service dynamodb.DynamoDB, user User, tableName string) error {
	av, err := dynamodbattribute.MarshalMap(user)

	fmt.Print(av)
	if err != nil {
		return err
	}

	_, err = service.PutItem(&dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      av,
	})

	if err != nil {
		return err
	}

	return nil
}
