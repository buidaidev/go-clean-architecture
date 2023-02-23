package main

import (
	"log"

	"github.com/joho/godotenv"
)

func InitialEnvironment() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
