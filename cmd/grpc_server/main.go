package main

import (
	"context"
	"log"
	"week/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
