package main

import (
	"log"
	"os"

	"github.com/ikovic/notion-todo-daily/cmd"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	log.Println(os.Getenv("AUTH_TOKEN"))

	cmd.Execute()
}
