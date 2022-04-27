package infra

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/italorfeitosa/url-shortner-mvp/libs/event"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/domain"
)

var topics map[string]string = map[string]string{
	domain.ShortURLCreatedTopic: "arn:aws:sns:us-east-2:000000000000:url_management_short_url_topic.fifo",
}

type SNSBroker struct {
	s *sns.Client
}

func NewSNSBroker() *SNSBroker {
	s := NewSNS()
	return &SNSBroker{s}
}

func (s *SNSBroker) Publish(ctx context.Context, e event.Event) error {
	log.Println("publish event ", e.Topic)
	message, err := e.ToJSON()

	if err != nil {
		return err
	}

	topicArn, ok := topics[e.Topic]

	if !ok {
		return errors.New("topic arn not found")
	}

	_, err = s.s.Publish(ctx, &sns.PublishInput{
		Message:        aws.String(message),
		TopicArn:       aws.String(topicArn),
		MessageGroupId: aws.String(fmt.Sprint(e.Kind(), "-", e.AggregateID)),
	})

	if err != nil {
		return err
	}

	log.Println("event published ", e.Topic)

	return nil
}
