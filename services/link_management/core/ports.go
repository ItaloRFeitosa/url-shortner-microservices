package core

import (
	"context"
)

type CreateLinkGateway interface {
	Create(ctx context.Context, link *Link) error
}

type ListLinksGateway interface {
	List(ctx context.Context, workspaceID string) ([]Link, error)
}

type LinkRepository interface {
	CreateLinkGateway
	ListLinksGateway
}

type CreateLinkInput struct {
	WorkspaceID string
	RedirectTo  string
	Title       string
	Description string
}

type CreateLinkUseCase interface {
	Create(c context.Context, input CreateLinkInput) (*Link, error)
}

type LinkUseCase interface {
	CreateLinkUseCase
}

type Encoder interface {
	EncodeSlug(link *Link) error
}
