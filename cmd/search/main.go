package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/letstrygo/templates/internal/database"
)

func main() {
	var (
		search string
	)

	if len(os.Args) > 1 {
		search = os.Args[1]
	}

	tmpls, err := database.ListTemplates(database.ListTemplatesArg{
		Search: search,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, tmpl := range tmpls {
		data, err := json.Marshal(tmpl)
		if err == nil {
			println(string(data))
		}
	}
}
