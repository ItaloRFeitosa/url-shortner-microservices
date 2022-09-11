package service

import (
	"context"

	"github.com/italorfeitosa/url-shortner-mvp/services/link_redirect/domain"
	"github.com/italorfeitosa/url-shortner-mvp/services/link_redirect/dto"
)

type ServiceImpl struct {
	r Repository
	e Encoder
}

func NewService(r Repository, e Encoder) *ServiceImpl {
	return &ServiceImpl{r, e}
}

func (s *ServiceImpl) AddShortURL(ctx context.Context, input dto.AddShortURLDTO) error {
	url := domain.NewShortURL(input)

	err := s.r.Insert(ctx, url)

	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) GetRedirectURL(ctx context.Context, slug string) (string, error) {
	id, err := s.e.Decode(slug)

	if err != nil {
		return "", err
	}

	url, err := s.r.FindRedirectURL(ctx, id)

	if err != nil {
		return "", err
	}

	// TODO: dispatch event

	return url, nil
}
