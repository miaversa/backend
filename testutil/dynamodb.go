package testutil

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoDB() *dynamodb.DynamoDB {
	debug := false
	sess, err := session.NewSession(&aws.Config{Region: aws.String("sa-east-1"), Endpoint: aws.String("http://localhost:8000")})
	if err != nil {
		panic(err)
	}
	if debug {
		return dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	}
	return dynamodb.New(sess, aws.NewConfig())
}
