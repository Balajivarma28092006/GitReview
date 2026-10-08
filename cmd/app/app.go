package app

import (
	"fmt"

	"github.com/Balajivarma28092006/GitReview/internal/database"
	"github.com/Balajivarma28092006/GitReview/internal/users"
	"github.com/gin-gonic/gin"
)

func RunApp() error {
	db, err := database.New()
	if err != nil {
		return fmt.Errorf("something went wrong: %w", err)
	}
	defer db.Close()

	repo := users.NewRepository(db.Pool())
	service := users.NewService(repo)
	handler := users.NewUserHandler(service)

	router := gin.Default()
	router.POST("/api/users", handler.CreateUser)
	if err := router.Run(); err != nil {
		return fmt.Errorf("somethings wrong: %w", err)
	}
	return nil
}
