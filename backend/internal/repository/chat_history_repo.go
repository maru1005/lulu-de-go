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
	INSERT INTO lulu.chat_history (user_id, category, role, content)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
	`

	return r.db.QueryRowContext(ctx, query,
		c.UserID, c.Category, c.Role, c.Content).Scan(&c.ID, &c.CreatedAt)
}

func (r *ChatHistoryRepository) ListByUser(ctx context.Context, userID string) ([]*model.ChatHistory, error) {
	query := `
		SELECT id, user_id, category, role, content, created_at
		FROM lulu.chat_history
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []*model.ChatHistory
	for rows.Next() {
		var c model.ChatHistory
		if err := rows.Scan(
			&c.ID, &c.UserID, &c.Category, &c.Role, &c.Content, &c.CreatedAt,
		); err != nil {
			return nil, err
		}
		histories = append(histories, &c)
	}

	return histories, rows.Err()
}

func (r *ChatHistoryRepository) TrimToLimit(ctx context.Context, userID string, limit int) error {
	query := `
		DELETE FROM lulu.chat_history
		WHERE user_id = $1
		AND id NOT IN (
			SELECT id FROM lulu.chat_history
			WHERE user_id = $1
			ORDER BY created_at DESC
			LIMIT $2
		)
	`

	_, err := r.db.ExecContext(ctx, query, userID, limit)
	return err
}
