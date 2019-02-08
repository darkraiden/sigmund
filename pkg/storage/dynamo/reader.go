package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// ReadFromTable executes a Select query on the DynamoDB
// database returning the output in an Item format
func (client *Client) ReadFromTable(table string) (*Item, error) {
	var item Item
	result, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String("0"),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Error getting item from DynamoDB: %v", err)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal record: %v", err)
	}

	return &item, nil
}
