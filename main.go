package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// import .env file
	// capture the err
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

}
