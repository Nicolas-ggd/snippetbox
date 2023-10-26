package mocks

import (
	"time"

	"snippetbox.nicolas.net/internal/models"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "Example1",
	Content: "Example2",
	Expires: &time.Time{},
	Created: &time.Time{},
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title string, content string, created time.Time, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
