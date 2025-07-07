package main

import (
	"log"

	"github.com/letstrygo/templates/internal/database"
)

func main() {
	err := database.Migrate()
	if err != nil {
		log.Fatal(err)
	}
}
