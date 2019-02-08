package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// WriteToTable executes an Update query on the DynamoDB
// updating the provided key with the isLow bool
func (client *Client) WriteToTable(table, key string, isLow bool) error {
	input := dbUpdateItemInput(table, key, isLow)

	_, err := client.UpdateItem(input)

	if err != nil {
		return fmt.Errorf("Something went wrong while running the Update query: %v", err)
	}

	return nil
}

func dbUpdateItemInput(table, key string, isLow bool) *dynamodb.UpdateItemInput {
	return &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":metric": {
				BOOL: aws.Bool(isLow),
			},
		},
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String("0"),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(fmt.Sprintf("SET %s = :metric", key)),
	}
}
