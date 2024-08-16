package main

import (
	"github.com/joho/godotenv"
	"github.com/melnk300/medodsTest/internal/app"
	"log"
	"net/http"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	log.Println("Server started on port 3000")
	err := http.ListenAndServe(":3000", app.Server())
	if err != nil {
		panic("Server incorrect")
	}
}
