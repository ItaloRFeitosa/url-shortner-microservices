package database

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/italorfeitosa/url-shortner-mvp/services/link_management/core"
	"github.com/jmoiron/sqlx"
)

func TestLinkRepository_Create(t *testing.T) {
	db := NewDB()
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx  context.Context
		link *core.Link
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should add new link record to database",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				link: core.CreateLink(core.CreateLinkInput{
					WorkspaceID: "wid",
					RedirectTo:  "https://google.com",
					Title:       "Google",
					Description: "Google Page",
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &LinkRepository{
				db: tt.fields.db,
			}
			if err := r.Create(tt.args.ctx, tt.args.link); err != nil {
				t.Errorf("LinkRepository.Create() error = %v", err)
			}
			if tt.args.link.ID == 0 {
				t.Errorf("after LinkRepository.Create() expect link.ID to be defined. got link.ID = %v", tt.args.link.ID)
			}
			gotLink, err := r.FindByID(tt.args.ctx, tt.args.link.ID)
			if err != nil {
				t.Errorf("LinkRepository.FindByID() error = %v", err)
			}
			if !cmp.Equal(*tt.args.link, gotLink) {
				t.Errorf("after LinkRepository.FindByID() expect find same link. diff = %v", cmp.Diff(*tt.args.link, gotLink))
			}
		})
	}
}
