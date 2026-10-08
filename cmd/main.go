package main

import (
	"fmt"

	"github.com/Balajivarma28092006/GitReview/internal/database"
)

func main() {
	service, err := database.New()
	if err != nil {
		fmt.Printf("something went wrong: %v", err)
	}
	stat := service.Pool().Stat()
	fmt.Print(stat)
}
