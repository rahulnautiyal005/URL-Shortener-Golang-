package database

import (
	"database/sql"
	"fmt"
	"url-shortener/models"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(connStr string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Create table if not exists
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		long_url TEXT NOT NULL,
		short_code TEXT UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		expires_at TIMESTAMP,
		click_count INT DEFAULT 0
	);`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) GetNextID() (int64, error) {
	var id int64
	err := s.db.QueryRow("SELECT nextval(pg_get_serial_sequence('urls', 'id'))").Scan(&id)
	if err != nil {
		// If serial sequence fails, we can just use the count or another method
		// But for a simple project, we'll assume the sequence works.
		return 0, err
	}
	return id, nil
}

func (s *PostgresStore) Save(url *models.URL) error {
	query := `INSERT INTO urls (id, long_url, short_code, created_at, expires_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(query, url.ID, url.LongURL, url.ShortCode, url.CreatedAt, url.ExpiresAt)
	return err
}

func (s *PostgresStore) GetByCode(code string) (*models.URL, error) {
	var url models.URL
	query := `SELECT id, long_url, short_code, created_at, expires_at, click_count FROM urls WHERE short_code = $1`
	err := s.db.QueryRow(query, code).Scan(&url.ID, &url.LongURL, &url.ShortCode, &url.CreatedAt, &url.ExpiresAt, &url.ClickCount)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (s *PostgresStore) IncrementClick(code string) error {
	query := `UPDATE urls SET click_count = click_count + 1 WHERE short_code = $1`
	_, err := s.db.Exec(query, code)
	return err
}
