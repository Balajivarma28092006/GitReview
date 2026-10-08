package users

import (
	"context"
	"errors"
	"regexp"
	"time"
	"uuid"
)

var InvalidEmail = errors.New("Email is not valid")
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9.!#\$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

type UserService interface {
	AddUser(ctx context.Context, reqUser CreateRequestUser) (*User, error)
}

type Services struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Services {
	return &Services{
		repo: repo,
	}
}

func (service *Services) AddUser(ctx context.Context, reqUser CreateRequestUser) (*User, error) {
	var id uuid.UUID

	id = uuid.New()
	created_at := time.Now()

	if !emailRegex.MatchString(reqUser.Email) {
		return nil, InvalidEmail
	}

	user := &User{
		ID:        id,
		UserName:  reqUser.UserName,
		Email:     reqUser.Email,
		CreatedAt: created_at,
	}

	// 3 seconds time out
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user, err := service.repo.Insert(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, err
}
