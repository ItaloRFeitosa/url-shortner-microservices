package infra

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/italorfeitosa/url-shortner-mvp/libs/event"
)

type SQSBroker struct {
	s        *sqs.Client
	handlers map[string][]event.EventHandler
	queueUrl string
}

func NewSQSBroker() *SQSBroker {
	s := NewSQS()

	handlers := map[string][]event.EventHandler{}
	queueUrl := "http://localhost:4566/000000000000/url_redirect_short_url_queue.fifo"
	return &SQSBroker{s, handlers, queueUrl}
}

func (b *SQSBroker) Subscribe(topic string, handler event.EventHandler) error {
	if _, ok := b.handlers[topic]; !ok {
		b.handlers[topic] = []event.EventHandler{handler}

		return nil
	}

	b.handlers[topic] = append(b.handlers[topic], handler)

	return nil
}

func (b *SQSBroker) Init(ctx context.Context) {

	for {
		output, err := b.s.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(b.queueUrl),
			MaxNumberOfMessages: 10,
			WaitTimeSeconds:     10,
		})

		if err != nil {
			log.Printf("failed to fetch sqs message %v", err)
			continue
		}

		for _, message := range output.Messages {
			go b.handleMessage(ctx, message)
		}
	}
}

func (b *SQSBroker) handleMessage(ctx context.Context, sqsMessage types.Message) {
	message, err := GetBodyMessage(*sqsMessage.Body)

	if err != nil {
		log.Printf("failed to handle sqs message %v", err)
		return
	}

	e := event.Event{}

	err = e.BindJSON(message)

	if err != nil {
		log.Printf("failed to handle sqs message %v", err)
		return
	}

	handlers, ok := b.handlers[e.Topic]

	if !ok {
		log.Printf("not found handler to event.Topic %v", e.Topic)
		return
	}

	for _, h := range handlers {
		go b.handleEvent(ctx, h, e)
	}

	b.s.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(b.queueUrl),
		ReceiptHandle: sqsMessage.ReceiptHandle,
	})
}

func (b *SQSBroker) handleEvent(ctx context.Context, h event.EventHandler, e event.Event) {
	err := h.Handle(ctx, e)

	if err != nil {
		b.handleEventError(err)
	}
}

func (b *SQSBroker) handleEventError(err error) {
	if eventError, ok := err.(*event.EventError); ok {
		log.Printf("failed to handle event %v", eventError)
		return
	}

	log.Printf("failed to handle event %v", err)
}

func GetBodyMessage(v string) (string, error) {
	j := map[string]any{}
	err := json.Unmarshal([]byte(v), &j)

	if err != nil {
		return "", err
	}

	message, ok := j["Message"].(string)

	if !ok {
		return "", errors.New("invalid body message")
	}

	return message, nil
}
