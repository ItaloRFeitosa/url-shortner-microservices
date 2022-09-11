package aws_helper

import (
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func NewSNS() *sns.Client {
	cfg := NewConfig()

	return sns.NewFromConfig(cfg)
}
