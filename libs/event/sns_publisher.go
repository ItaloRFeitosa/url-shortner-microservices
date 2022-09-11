package event

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/italorfeitosa/url-shortner-mvp/libs/aws_helper"
)

type SNSPublisherConfig struct {
	Topics map[string]string
}

type SNSPublisher struct {
	s      *sns.Client
	config SNSPublisherConfig
}

func NewSNSPublisher(config SNSPublisherConfig) *SNSPublisher {
	s := aws_helper.NewSNS()
	return &SNSPublisher{s, config}
}

func (s *SNSPublisher) Publish(ctx context.Context, e Event) error {
	log.Println("publish event ", e.Topic)
	message, err := e.ToJSON()

	if err != nil {
		return err
	}

	topicArn, ok := s.config.Topics[e.Topic]

	if !ok {
		return errors.New("topic arn not found")
	}

	_, err = s.s.Publish(ctx, &sns.PublishInput{
		Message:        aws.String(message),
		TopicArn:       aws.String(topicArn),
		MessageGroupId: aws.String(fmt.Sprint(e.Kind(), "@", e.AggregateID)),
	})

	if err != nil {
		return err
	}

	log.Println("event published ", e.Topic)

	return nil
}
