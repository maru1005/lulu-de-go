package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/maru1005/lulu-de-go/internal/model"
)

type QuestionRepository interface {
	GetNext(ctx context.Context, userID, category string) (*model.Question, error)
}

var ErrPoolEmpty = errors.New("question pool is empty")

type QuestionService struct {
	questionRepo QuestionRepository
}

func NewQuestionService(questionRepo QuestionRepository) *QuestionService {
	return &QuestionService{questionRepo: questionRepo}
}

func (s *QuestionService) GetNext(ctx context.Context, userID, category, mode string) (*model.Question, error) {
	switch mode {
	case "new":
		q, err := s.questionRepo.GetNext(ctx, userID, category)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrPoolEmpty
			}
			return nil, err
		}
		return q, nil
	default:
		return nil, fmt.Errorf("unsupported mode: %q", mode)
	}
}
