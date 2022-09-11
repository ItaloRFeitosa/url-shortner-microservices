package core

import (
	"strconv"

	"github.com/italorfeitosa/url-shortner-mvp/libs/event"
)

const LinkCreatedTopic string = "link.created"

func LinkCreatedEvent(link Link) event.Event {
	return event.New(event.NewEventInput{
		AggregateID: strconv.FormatUint(link.ID, 10),
		Topic:       LinkCreatedTopic,
		Data:        link,
	})
}
