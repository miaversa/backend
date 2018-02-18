package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ensureCartsTable(cli *dynamodb.DynamoDB) error {
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
