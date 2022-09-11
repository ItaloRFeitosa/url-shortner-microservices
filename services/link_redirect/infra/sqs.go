package infra

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/italorfeitosa/url-shortner-mvp/libs/aws_helper"
)

func NewSQS() *sqs.Client {
	cfg := aws_helper.NewConfig()

	return sqs.NewFromConfig(cfg)
}
