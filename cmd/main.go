package main

import (
	"github.com/Balajivarma28092006/GitReview/cmd/app"
)

func main() {
	err := app.RunApp()
	if err != nil {
		panic(err)
	}
}
