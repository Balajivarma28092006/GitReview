package posts

import (
	"time"

	"github.com/google/uuid"
)

type PostID = uuid.UUID

type Posts struct {
	ID        PostID    `json:"id"`
	Content   string    `json:"content"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
