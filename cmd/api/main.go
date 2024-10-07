package main

import (
	"log"

	"github.com/romacardozx/online-payment-platform/internal/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}
}
