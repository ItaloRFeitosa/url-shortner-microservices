package event

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Event struct {
	AggregateID string    `json:"aggregate_id"`
	Topic       string    `json:"topic"`
	Data        any       `json:"data"`
	CreatedAt   time.Time `json:"created_at"`
}

func (e Event) ToJSON() (string, error) {
	b, err := json.Marshal(e)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (e *Event) BindJSON(j string) error {
	return json.Unmarshal([]byte(j), e)
}

func (e Event) Kind() string {
	return strings.Split(e.Topic, ".")[0]
}

func NewTopic(topic string) func(string, any) Event {
	return func(aggregateID string, data any) Event {
		return Event{
			AggregateID: aggregateID,
			Topic:       topic,
			Data:        data,
			CreatedAt:   time.Now().UTC(),
		}
	}
}

type Publisher interface {
	Publish(ctx context.Context, event Event) error
}

type Subscriber interface {
	Subscribe(topic string, handler EventHandler) error
}

type EventHandler interface {
	Name() string
	Handle(ctx context.Context, event Event) error
}

type EventError struct {
	Event        Event
	EventHandler EventHandler
	err          error
}

func NewEventError(err error, e Event, h EventHandler) error {
	return &EventError{e, h, err}
}

func (e *EventError) Error() string {
	return fmt.Sprint("event error ", e.err, " on handler ", e.EventHandler.Name())
}
