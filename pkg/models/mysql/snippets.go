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

	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES (?,?, UTC _TIMESTAMP(), DATE_ADD(UT_TIMESTAMPP(), INTERVAL ? DAY))
	`
	result, err := m.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	//retrieve id of the newly created snippet from the table

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// return a specific snippet from DB
func (m *SnippetModel) Get(id int) (int, error) {
	return 0, nil
}

// return 10 mostly recently used snippents
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
