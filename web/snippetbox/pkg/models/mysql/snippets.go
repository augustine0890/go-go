package mysql

import (
	"database/sql"
	"snippetbox/pkg/models"
)

// wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &SnippetModel{
		DB: db,
	}
}

type Repository interface {
	Insert(title, content, expries string) (int, error)
	Get(id int) (*models.Snippet, error)
	Latest() ([]*models.Snippet, error)
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
