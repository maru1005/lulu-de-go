package model

import "time"

type Question struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	Category     string    `db:"category"`
	QuestionText string    `db:"question_text"`
	Choices      []string  `db:"choices"`
	CorrectIndex int       `db:"correct_index"`
	Hint         string    `db:"hint"`
	CreatedAt    time.Time `db:"created_at"`
}
