package domain

import (
	"time"

	"github.com/italorfeitosa/url-shortner-mvp/services/link_redirect/dto"
)

type ShortURL struct {
	ID         int       `json:"id"`
	RedirectTo string    `json:"redirect_to"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewShortURL(dto dto.AddShortURLDTO) *ShortURL {
	return &ShortURL{
		ID:         dto.ID,
		RedirectTo: dto.RedirectTo,
		CreatedAt:  time.Now().UTC(),
	}
}
