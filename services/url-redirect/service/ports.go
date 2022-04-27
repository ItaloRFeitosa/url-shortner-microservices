package service

import (
	"context"

	"github.com/italorfeitosa/url-shortner-mvp/services/url-redirect/domain"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-redirect/dto"
)

type Repository interface {
	Insert(c context.Context, u *domain.ShortURL) error
	FindRedirectURL(c context.Context, ID int) (string, error)
}

type Service interface {
	AddShortURL(c context.Context, input dto.AddShortURLDTO) error
	GetRedirectURL(c context.Context, slug string) (string, error)
}

type Encoder interface {
	Decode(v string) (int, error)
}
