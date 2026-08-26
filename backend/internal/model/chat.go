package model

import "time"

type ChatHistory struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Category  *string   `db:"category"`
	Role      string    `db:"role"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
