package entities

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// User is a struct that implements the Item interface
type User struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

// GetPK returns the partition key for the user item
func (user User) GetPK() *string {
	primaryKey := fmt.Sprintf("USER#%s", *user.Email)
	return &primaryKey
}

// GetSK returns the sort key for the user item
func (user User) GetSK() *string {
	sortKey := fmt.Sprintf("USER#%s", *user.Email)
	return &sortKey
}

// Keys returns the DynamoDB keys for the user item
func (user User) Keys() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"PK": {S: user.GetPK()},
		"SK": {S: user.GetSK()},
	}
}

// ToItem returns the DynamoDB item
func (user User) ToItem() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"PK":    {S: user.GetPK()},
		"SK":    {S: user.GetSK()},
		"name":  {S: user.Name},
		"email": {S: user.Email},
	}
}

// FromItem returns the UserItem from the DynamoDB Item
func (user User) FromItem(item map[string]*dynamodb.AttributeValue) Item {
	name := *item["name"].S
	email := *item["email"].S

	return User{
		Name:  &name,
		Email: &email,
	}
}
