package service

import (
	"context"

	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/domain"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/dto"
)

type Repository interface {
	Insert(c context.Context, u *domain.ShortURL) error
	ListUserURLs(c context.Context, userID string) ([]domain.ShortURL, error)
}

type Service interface {
	ShortenURL(c context.Context, input dto.ShortenURLDTO) (*domain.ShortURL, error)
	ListUserURLs(c context.Context, userID string) ([]domain.ShortURL, error)
}

type Encoder interface {
	Encode(v int) (string, error)
}
