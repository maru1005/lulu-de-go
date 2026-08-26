package model

import "time"

type History struct {
	ID            string    `db:"id"`
	UserID        string    `db:"user_id"`
	Category      string    ` db:"category"`
	Mode          string    `db:"mode"`
	QuestionText  string    `db:"question_text"`
	SelectedIndex int       `db:"selected_index"`
	CorrectIndex  int       `db:"correct_index"`
	IsCorrect     bool      `db:"is_correct"`
	CreatedAt     time.Time `db:"created_at"`
}
