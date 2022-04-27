package event_handler

import (
	"context"
	"log"

	"github.com/italorfeitosa/url-shortner-mvp/libs/event"
	"github.com/italorfeitosa/url-shortner-mvp/libs/map_helper"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-redirect/dto"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-redirect/service"
)

type AddShortURLHandler struct {
	s service.Service
}

func NewAddShortURLHandler(s service.Service) *AddShortURLHandler {
	return &AddShortURLHandler{s}
}

func (a *AddShortURLHandler) Name() string {
	return "AddShortURLHandler"
}

func (a *AddShortURLHandler) Handle(ctx context.Context, e event.Event) error {
	d, err := map_helper.Decode[dto.AddShortURLDTO](e.Data)

	if err != nil {
		return event.NewEventError(err, e, a)
	}
	log.Println("adding url id: ", d.ID, "redirect_to: ", d.RedirectTo)

	err = a.s.AddShortURL(ctx, d)

	if err != nil {
		return event.NewEventError(err, e, a)
	}
	return nil
}
