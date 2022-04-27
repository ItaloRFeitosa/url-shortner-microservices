package aws_helper

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func NewConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("us-east-2"), config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{URL: "http://localhost:4566", SigningRegion: region}, nil
		})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "fake", SecretAccessKey: "fake", SessionToken: "fake",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}))
	if err != nil {
		panic(err)
	}

	return cfg
}
