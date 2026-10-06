package users

import (
	"time"
	"uuid"
)

type UserID uuid.UUID

type User struct {
	ID        UserID    `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// without any uuid, the uuid will be generated afterwords
type CreateRequestUser struct {
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
