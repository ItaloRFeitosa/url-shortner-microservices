package domain

import (
	"time"

	"github.com/italorfeitosa/url-shortner-mvp/libs/event"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/dto"
)

const ShortURLCreatedTopic string = "short_url.created"

var NewShortUrlCreatedEvent func(string, any) event.Event = event.NewTopic(ShortURLCreatedTopic)

type ShortURL struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	Slug        string    `json:"slug"`
	RedirectTo  string    `json:"redirect_to"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewShortURL(dto dto.ShortenURLDTO) *ShortURL {
	return &ShortURL{
		UserID:      dto.UserID,
		RedirectTo:  dto.RedirectTo,
		Title:       dto.Title,
		Description: dto.Description,
		CreatedAt:   time.Now().UTC(),
	}
}
