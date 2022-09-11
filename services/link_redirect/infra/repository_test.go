package infra

import (
	"context"
	"testing"

	"github.com/italorfeitosa/url-shortner-mvp/services/link_redirect/domain"
	"github.com/italorfeitosa/url-shortner-mvp/services/link_redirect/dto"
)

func TestDynamoDBRepository_Insert(t *testing.T) {
	type args struct {
		ctx context.Context
		url *domain.ShortURL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "should insert ShortUrl in db",
			args: args{
				ctx: context.Background(),
				url: domain.NewShortURL(dto.AddShortURLDTO{
					RedirectTo: "https://instagram.com/italo_feittosa",
					ID:         1, // Bp69
				}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDynamoDBRepository()
			if err := d.Insert(tt.args.ctx, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("DynamoDBRepository.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
