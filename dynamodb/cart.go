package dynamodb

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/miaversa/backend/cart"
)

type dynamoDBCartStorage struct {
	cli *dynamodb.DynamoDB
}

func NewCartStorage(cli *dynamodb.DynamoDB) (*dynamoDBCartStorage, error) {
	err := ensureTable(cli)
	if err != nil {
		return nil, err
	}
	return &dynamoDBCartStorage{cli: cli}, nil
}

func ensureTable(cli *dynamodb.DynamoDB) error {
	result, err := cli.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return err
	}
	var found bool
	for _, v := range result.TableNames {
		defV := *v
		if defV == "carts" {
			found = true
		}
	}
	if !found {
		input := &dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("key"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("key"),
					KeyType:       aws.String("HASH"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(10),
				WriteCapacityUnits: aws.Int64(10),
			},
			TableName: aws.String("carts"),
		}
		_, err = cli.CreateTable(input)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *dynamoDBCartStorage) GetCart(id string) (cart.Cart, error) {
	result, err := s.cli.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("carts"),
		Key: map[string]*dynamodb.AttributeValue{
			"key": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return cart.New(id), err
	}
	if len(result.Item) == 0 {
		return cart.New(id), errors.New("item not found")
	}
	cartItem := cart.Cart{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &cartItem)
	if err != nil {
		return cart.New(id), err
	}
	return cartItem, nil
}

func (s *dynamoDBCartStorage) SaveCart(c cart.Cart) error {
	av, err := dynamodbattribute.MarshalMap(c)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("carts"),
	}
	_, err = s.cli.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}

func (s *dynamoDBCartStorage) DropCart(id string) error {
	return nil
}
