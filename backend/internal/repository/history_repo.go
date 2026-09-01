package repository

import (
	"context"
	"database/sql"

	"github.com/maru1005/lulu-de-go/internal/model"
)

type HistoryRepository struct {
	db *sql.DB
}

func NewHistoryRepository(db *sql.DB) *HistoryRepository {
	return &HistoryRepository{db: db}
}

func (r *HistoryRepository) Create(ctx context.Context, h *model.History) error {
	query := `
	INSERT INTO lulu.history (user_id, category, mode, question_text, selected_index, correct_index, is_correct)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, created_at
	`

	return r.db.QueryRowContext(ctx, query,
		h.UserID, h.Category, h.Mode, h.QuestionText, h.SelectedIndex, h.CorrectIndex).Scan(&h.ID, &h.CreatedAt)
}

func (r *HistoryRepository) ListByUser(ctx context.Context, userID string) ([]*model.History, error) {
	query := `
	SELECT id, user_id, category, mode, question_text, selected_index, correct_index, is_correct, created_at
	FROM lulu.history
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close() //終了後閉じる

	var histories []*model.History
	for rows.Next() { // 1行づつループで取り出す
		var h model.History
		if err := rows.Scan(
			&h.ID, &h.UserID, &h.Category, &h.Mode, &h.QuestionText, &h.SelectedIndex, &h.CorrectIndex, &h.IsCorrect, &h.CreatedAt,
		); err != nil {
			return nil, err
		}
		histories = append(histories, &h)
	}

	return histories, rows.Err() // ループを抜けたあとにチェック
}
