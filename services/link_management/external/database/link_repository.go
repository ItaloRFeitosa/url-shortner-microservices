package database

import (
	"context"
	"time"

	"github.com/italorfeitosa/url-shortner-mvp/services/link_management/core"
	"github.com/jmoiron/sqlx"
)

type LinkRepository struct {
	db *sqlx.DB
}

const (
	InsertIntoLinks         = "INSERT INTO links (Title, Description, RedirectTo, WorkspaceID, CreatedAt) VALUES (?, ?, ?, ?, ?)"
	SelectFromLinksPaginate = "SELECT * FROM links WHERE WorkspaceID=? ORDER BY created_at DESC LIMIT ? OFFSET ?"
	SelectFromLinksWhereID  = "SELECT * FROM links WHERE ID=? LIMIT 1"
)

func NewLinkRepository(db *sqlx.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) Create(ctx context.Context, link *core.Link) (err error) {
	result, err := r.db.ExecContext(
		ctx,
		InsertIntoLinks,
		link.Title,
		link.Description,
		link.RedirectTo,
		link.WorkspaceID,
		link.CreatedAt.Truncate(1*time.Second).Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return err
	}
	linkID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	link.ID = uint64(linkID)
	return nil
}

func (r *LinkRepository) FindByID(ctx context.Context, ID uint64) (core.Link, error) {
	link := core.Link{}
	if err := r.db.GetContext(ctx, &link, SelectFromLinksWhereID, ID); err != nil {
		return link, err
	}
	return link, nil
}

func (r *LinkRepository) List(ctx context.Context, workspaceID string) ([]core.Link, error) {
	links := []core.Link{}
	if err := r.db.SelectContext(ctx, &links, SelectFromLinksPaginate, workspaceID, 10, 0); err != nil {
		return links, err
	}
	return links, nil
}
