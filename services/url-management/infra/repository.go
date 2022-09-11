package infra

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/italorfeitosa/url-shortner-mvp/services/url-management/domain"
)

type MySQLRepository struct {
	*sql.DB
}

func NewMySQLRepository() *MySQLRepository {
	db := CreateConnection()
	return &MySQLRepository{db}
}

func (m *MySQLRepository) Insert(ctx context.Context, url *domain.ShortURL) error {
	stmt, err := m.PrepareContext(
		ctx,
		`INSERT INTO short_urls (Title, Description, RedirectTo, UserID, CreatedAt) VALUES (?, ?, ?, ?, ?)`,
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(
		ctx,
		url.Title,
		url.Description,
		url.RedirectTo,
		url.UserID,
		url.CreatedAt.Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return err
	}

	url.ID = int(id)

	return nil
}

func (m *MySQLRepository) ListUserURLs(ctx context.Context, userID string) ([]domain.ShortURL, error) {
	urls := []domain.ShortURL{}
	stmt, err := m.PrepareContext(
		ctx,
		`SELECT ID, RedirectTo, CreatedAt FROM short_urls WHERE UserID=?`,
	)

	if err != nil {
		return urls, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userID)

	if err != nil {
		return urls, err
	}

	defer rows.Close()

	for rows.Next() {
		url := domain.ShortURL{}
		var date string
		err := rows.Scan(&url.ID, &url.RedirectTo, &date)
		if err != nil {
			return urls, err
		}

		url.CreatedAt, err = time.Parse("2006-01-02 15:04:05", date)

		if err != nil {
			return urls, err
		}

		urls = append(urls, url)
	}
	return urls, nil
}

func (m *MySQLRepository) Disconnect() {
	err := m.Close()

	if err != nil {
		log.Fatal(err)
	}
}
