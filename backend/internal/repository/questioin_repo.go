package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/maru1005/lulu-de-go/internal/model"
)

type QuestionRepository struct {
	db *sql.DB
}

// 問題を作る
func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) Create(ctx context.Context, q *model.Question) error {
	choicesJSON, err := json.Marshal(q.Choices)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO lulu.questions (user_id, category, question_text, choices, correct_index, hint)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, created_at
	`

	return r.db.QueryRowContext(ctx, query,
		q.UserID, q.Category, q.QuestionText, choicesJSON, q.CorrectIndex, q.Hint,
	).Scan(&q.ID, &q.CreatedAt)
}

// 出題　古いものから
func (r *QuestionRepository) GetNext(ctx context.Context, userId, category string) (*model.Question, error) {

	query := `
	SELECT id, user_id, category, question_text, choices, correct_index, hint, created_at
	FROM lulu.questions
	WHERE user_id = $1 AND categoty = $2
	ORDER BY created_at ASC
	LIMIT 1
	`

	var q model.Question
	var choicesJSON []byte

	err := r.db.QueryRowContext(ctx, query, userId, category).Scan(
		&q.ID, &q.UserID, &q.Category, &q.QuestionText, &choicesJSON, &q.CorrectIndex, &q.Hint, &q.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(choicesJSON, &q.Choices); err != nil {
		return nil, err
	}

	return &q, nil
}

// 出題後プールから削除
func (r *QuestionRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM lulu.questions WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// プール件数
func (r *QuestionRepository) Count(ctx context.Context, userID, category string) (int, error) {
	query := `
	SELECT COUNT(*) FROM lulu.questions
	WHERE user_id = $1 AND category = $2
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, userID, category).Scan(&count)
	return count, err
}
