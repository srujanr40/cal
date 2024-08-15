package db

import (
    "context"
    "log"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func NewDynamoClient(cfg aws.Config) *dynamodb.Client {
    return dynamodb.NewFromConfig(cfg)
}

func CreateItem(client *dynamodb.Client, tableName string, item interface{}) error {
    // Implement the logic to put an item in the DynamoDB table
    // Use client.PutItem with the correct input
    return nil
}

func GetItems(client *dynamodb.Client, tableName string) ([]map[string]types.AttributeValue, error) {
    // Implement the logic to get items from the DynamoDB table
    // Use client.Scan with the correct input
    return nil, nil
}