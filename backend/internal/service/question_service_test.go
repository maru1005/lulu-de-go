package service

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/maru1005/lulu-de-go/internal/model"
)

// 本物のDBに繋げない　テスト用
// QuestionRepository interfaceが求めるGetNextさえ持っていれば、
// それだけでこのinterfaceを満たしたことになる
type fakeQuestionRepo struct {
	question *model.Question
	err      error
}

func (f *fakeQuestionRepo) GetNext(ctx context.Context, userID, category string) (*model.Question, error) {
	return f.question, f.err
}

func TestQuestionService_GetNext(t *testing.T) {
	tests := []struct {
		name      string
		mode      string
		repo      *fakeQuestionRepo
		wantErr   bool
		wantErrIs error // 特定のerrorを期待する時だけ
	}{
		{
			name:    "modeがnewで正常に取得できる",
			mode:    "new",
			repo:    &fakeQuestionRepo{question: &model.Question{ID: "q1", Category: "basic"}},
			wantErr: false,
		},
		{
			name:      "modeがnewでプールが空ならErrPoolEmpty",
			mode:      "new",
			repo:      &fakeQuestionRepo{err: sql.ErrNoRows},
			wantErr:   true,
			wantErrIs: ErrPoolEmpty,
		},
		{
			name:    "未対応のmodeはエラー",
			mode:    "unknown",
			repo:    &fakeQuestionRepo{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewQuestionService(tt.repo)
			_, err := s.GetNext(context.Background(), "user1", "basic", tt.mode)

			if tt.wantErr && err == nil {
				t.Fatal("エラーを期待したがnilだった")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("エラーは期待してなかったが発生した: %v", err)
			}
			if tt.wantErrIs != nil && !errors.Is(err, tt.wantErrIs) {
				t.Errorf("got err = %v, want %v", err, tt.wantErrIs)
			}
		})
	}
}
