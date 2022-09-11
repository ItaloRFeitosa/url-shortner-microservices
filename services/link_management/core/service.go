package core

import (
	"context"

	"github.com/italorfeitosa/url-shortner-mvp/libs/event"
)

type LinkService struct {
	repo    LinkRepository
	encoder Encoder
	broker  event.Publisher
}

func NewLinkService(repo LinkRepository, encoder Encoder, broker event.Publisher) *LinkService {
	return &LinkService{repo, encoder, broker}
}
func (l *LinkService) Create(ctx context.Context, in CreateLinkInput) (*Link, error) {
	link := CreateLink(in)
	if err := l.repo.Create(ctx, link); err != nil {
		return nil, err
	}
	if err := l.encoder.EncodeSlug(link); err != nil {
		return nil, err
	}
	if err := l.broker.Publish(ctx, LinkCreatedEvent(*link)); err != nil {
		return nil, err
	}
	return link, nil
}
