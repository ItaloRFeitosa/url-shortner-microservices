package service

import (
	"context"
	"strconv"

	"github.com/italorfeitosa/url-shortner-mvp/libs/event"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/domain"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/dto"
)

type ServiceImpl struct {
	r Repository
	e Encoder
	p event.Publisher
}

func NewService(r Repository, e Encoder, p event.Publisher) *ServiceImpl {
	return &ServiceImpl{r, e, p}
}

func (s *ServiceImpl) ShortenURL(ctx context.Context, input dto.ShortenURLDTO) (*domain.ShortURL, error) {
	url := domain.NewShortURL(input)

	err := s.r.Insert(ctx, url)
	if err != nil {
		return nil, err
	}

	encodedId, err := s.e.Encode(url.ID)
	if err != nil {
		return nil, err
	}

	url.Slug = encodedId

	err = s.p.Publish(ctx, domain.NewShortUrlCreatedEvent(strconv.Itoa(url.ID), *url))

	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *ServiceImpl) ListUserURLs(ctx context.Context, userID string) ([]domain.ShortURL, error) {
	return s.r.ListUserURLs(ctx, userID)
}
