package core

import (
	"time"
)

type Link struct {
	ID          uint64    `json:"id" db:"ID"`
	Title       string    `json:"title" db:"Title"`
	Description string    `json:"description" db:"Description"`
	WorkspaceID string    `json:"workspace_id" db:"WorkspaceID"`
	Slug        string    `json:"slug"`
	RedirectTo  string    `json:"redirect_to" db:"RedirectTo"`
	CreatedAt   time.Time `json:"created_at" db:"CreatedAt"`
}

func CreateLink(input CreateLinkInput) *Link {
	return &Link{
		WorkspaceID: input.WorkspaceID,
		RedirectTo:  input.RedirectTo,
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   time.Now().UTC(),
	}
}
