package model

import "time"

type WeakPoint struct {
	ID            string    `db:"id"`
	UserID        string    `db:"user_id"`
	Category      string    `db:"category"`
	QuestioinText string    `db:"question_text"`
	Choices       []string  `db:"choices"`
	CorrectIndex  int       `db:"correct_index"`
	Explanation   string    `db:"explanation"`
	Embedding     []float32 `db:"ebbdding"`
	WrongCount    int       `db:"wrong_count"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
