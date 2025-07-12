package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/letstrygo/templates"
)

var (
	ErrInvalidArgumentCount error = errors.New("invalid argument count")
)

func main() {
	args := os.Args[1:]

	conn, err := templates.NewConnectionWithPath("database.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	switch args[0] {
	case "migrate":
		err = conn.Migrate()
		if err != nil {
			log.Fatal(err)
		}
	case "seed":
		err = conn.Migrate()
		if err != nil {
			log.Fatal(err)
		}

		err = conn.Seed()
		if err != nil {
			log.Fatal(err)
		}
	case "update":
		err = conn.Update()
		if err != nil {
			log.Fatal(err)
		}
	case "add":
		args = args[1:]

		if len(args) != 5 {
			log.Fatal(ErrInvalidArgumentCount)
		}

		err := conn.CreateTemplate(templates.CreateTemplate{
			Name:        args[0],
			Author:      args[1],
			AuthorURL:   args[2],
			CloneURL:    args[3],
			Description: args[4],
			IsOfficial:  false,
		})

		if err != nil {
			log.Fatal(err)
		}
	case "delete":
		args = args[1:]

		if len(args) != 1 {
			log.Fatal(ErrInvalidArgumentCount)
		}

		name := args[0]
		err := conn.DeleteTemplate(name)
		if err != nil {
			log.Fatal(err)
		}
	case "search":
		args = args[1:]

		var (
			search string
		)

		if len(args) > 0 {
			search = args[0]
		}

		conn, err := templates.NewConnectionWithPath("database.sqlite")
		if err != nil {
			log.Fatal(err)
		}

		tmpls, err := conn.ListTemplates(templates.ListTemplates{
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
}
