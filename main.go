package main

import (
	"context"
	"log"
	"os"

	"github.com/ikovic/notion-todo-daily/cmd"
	"github.com/ikovic/notion-todo-daily/internal/ctx"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	ctx := context.WithValue(context.Background(), ctx.ContextKey("AUTH_TOKEN"), os.Getenv("AUTH_TOKEN"))

	cmd.Execute(ctx)
}
