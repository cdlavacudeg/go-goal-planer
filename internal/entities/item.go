package entities

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Item is an abstract type for DynamoDB items
type Item interface {
	GetPK() *string                                         // Get the partition key
	GetSK() *string                                         // Get the sort key
	Keys() map[string]*dynamodb.AttributeValue              // Get the DynamoDB keys (PK and SK)
	ToItem() map[string]*dynamodb.AttributeValue            // Get the DynamoDB item
	FromItem(item map[string]*dynamodb.AttributeValue) Item // Get the Item from the DynamoDB item
}

// Crate a new item in dynamodb
func CreateItem(service dynamodb.DynamoDB, item Item, tableName string) error {
	userItem := item.ToItem()

	_, err := service.PutItem(&dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      userItem,
	})
	if err != nil {
		return err
	}

	return nil
}
