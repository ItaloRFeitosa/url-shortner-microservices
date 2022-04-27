package infra

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/italorfeitosa/url-shortner-mvp/libs/aws_helper"
)

func NewDynamoDB() *dynamodb.Client {
	cfg := aws_helper.NewConfig()

	return dynamodb.NewFromConfig(cfg)
}
