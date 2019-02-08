package dynamo

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Dynamo contains the DyanamoDB basic information
// needed to connect to the instance
type Dynamo struct {
	TableName string
	Region    string
	Key       string
}

// Client type is a DynamoDB client
type Client struct {
	*dynamodb.DynamoDB
}

// DB interface is a collection of most of
// the dyanmo package functions
type DB interface {
	NewSession() (*session.Session, error)
	NewClient() Client
	ReadFromTable(string) Item
	WriteToTable(string, string, bool) error
}

// New is the Package constructor that initialises
// the DynamoDB config
func New(tbl, region, key string) (*Dynamo, error) {
	config, err := checkConfig(tbl, region, key)
	if err != nil {
		return nil, fmt.Errorf("Initialisation error: %v", err)
	}

	return config, nil
}

// newSession creates a new AWS session to be used to
// connect to the DynamoDB instance
func newSession(region string) (*session.Session, error) {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// NewClient creates a DynamoDB client
func (dynamo *Dynamo) NewClient() (*Client, error) {
	var client Client

	sess, err := newSession(dynamo.Region)
	if err != nil {
		return nil, fmt.Errorf("Cannot open a new AWS session: %v", err)
	}

	client.DynamoDB = dynamodb.New(sess)
	return &client, nil
}

func checkConfig(tbl, region, key string) (*Dynamo, error) {
	switch {
	case tbl == "":
		return nil, errors.New("Table name cannot be empty")
	case region == "":
		return nil, errors.New("Region cannot be empty")
	case key == "":
		return nil, errors.New("Key cannot be empty")
	}
	return &Dynamo{
		TableName: tbl,
		Region:    region,
		Key:       key,
	}, nil
}
