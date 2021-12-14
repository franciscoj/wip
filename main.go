package main

import (
	"context"
	"log"

	"github.com/franciscoj/wip/internal/app"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("runtime: %v", err)
	}

	log.Println("Done!")
}

func run() error {
	ctx := context.Background()
	a, err := app.Load(ctx)
	if err != nil {
		return err
	}
	a.Print()

	return nil
}
