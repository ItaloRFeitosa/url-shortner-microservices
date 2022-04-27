package infra

import (
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/italorfeitosa/url-shortner-mvp/libs/aws_helper"
)

func NewSNS() *sns.Client {
	cfg := aws_helper.NewConfig()

	return sns.NewFromConfig(cfg)
}
