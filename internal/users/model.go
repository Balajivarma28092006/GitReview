package users

import (
	"time"
	"uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	UserName     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

// Public User is what users are allowed to see
type PublicUser struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) PublicUser() PublicUser {
	return PublicUser{
		ID:        u.ID,
		UserName:  u.UserName,
		CreatedAt: u.CreatedAt,
	}
}

// without any uuid, the uuid will be generated afterwords
type RegisterRequest struct {
	UserName string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"Token"`
	User  *User  `json:"user"`
}
