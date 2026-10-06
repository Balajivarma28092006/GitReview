package users

import (
	"time"
	"uuid"
)

type UserID uuid.UUID

type User struct {
	ID UserID `json:"id"`
	UserName string `json:"username"`
	Email string	`json:"email"`
	CreatedAt time.Time `json:"create_at"`
}


