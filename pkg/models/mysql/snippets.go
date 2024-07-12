package mysql

import (
	"database/sql"

	"github.com/flexarea/snippetbox/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.

type SnippetModel struct {
	DB *sql.DB
}

// insert a new snippet into DB
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// return a specific snippet from DB
func (m *SnippetModel) Get(id int) (int, error) {
	return 0, nil
}

// return 10 mostly recently used snippents
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
