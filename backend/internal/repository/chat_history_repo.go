package repository

import (
	"context"
	"database/sql"

	"github.com/maru1005/lulu-de-go/internal/model"
)

type ChatHistoryRepository struct {
	db *sql.DB
}

func NewChatHistoryRepository(db *sql.DB) *ChatHistoryRepository {
	return &ChatHistoryRepository{db: db}
}

func (r *ChatHistoryRepository) Create(ctx context.Context, c *model.ChatHistory) error {
	query := `
	INSERET INTO lulu.chat_history (user_id, category, role, content)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
	`

	return r.db.QueryRowContext(ctx, query,
		c.UserID, c.Category, c.Role, c.Content).Scan(&c.ID, &c.CreatedAt)
}
