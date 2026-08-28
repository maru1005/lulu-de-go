package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/maru1005/lulu-de-go/internal/model"
)

type WeakPointRepository struct {
	db *sql.DB
}

func NewWeakPointRepository(db *sql.DB) *WeakPointRepository {
	return &WeakPointRepository{db: db}
}

func (r *WeakPointRepository) Create(ctx context.Context, wp *model.WeakPoint) error {
	choicesJSON, err := json.Marshal(wp.Choices)
	if err != nil {
		return err
	}

	embeddingStr := vectorToString(wp.Embedding)

	query := `
	INSERT INTO lulu.weak_points (user_id, category, question_text, choices, correct_index, explanation, embedding, wrong_count)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowContext(ctx, query,
		wp.UserID, wp.Category, wp.QuestionText, choicesJSON, wp.CorrectIndex, wp.Explanation, embeddingStr, wp.WrongCount,
	).Scan(&wp.ID, &wp.CreatedAt, &wp.UpdatedAt)
}

func vectorToString(v []float32) string {
	strs := make([]string, len(v))
	for i, f := range v {
		strs[i] = fmt.Sprintf("%f", f)
	}
	return "[" + strings.Join(strs, ",") + "]"
}

func (r *WeakPointRepository) FindSimilar(ctx context.Context, userID, category string, embedding []float32, threshold float64) (*model.WeakPoint, error) {
	embeddingStr := vectorToString(embedding)

	query := `
	SELECT id, user_id, category, question_text, choices, correct_index, explanation, embedding, wrong_count, created_at, updated_at
	FROM lulu.weak_points
	WHERE user_id = $1 AND category = $2
	AND 1 - (embedding <=> $3 ) > $4
	ORDER BY embedding <=> $3
	LIMIT 1
	`

	var wp model.WeakPoint
	var choicesJSON []byte
	var embeddingRaw string

	err := r.db.QueryRowContext(ctx, query, userID, category, embeddingStr, threshold).Scan(
		&wp.ID, &wp.UserID, &wp.Category, &wp.QuestionText, &choicesJSON, &wp.CorrectIndex,
		&wp.Explanation, &embeddingRaw, &wp.WrongCount, &wp.CreatedAt, &wp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(choicesJSON, &wp.Choices); err != nil {
		return nil, err
	}
	return &wp, nil
}
