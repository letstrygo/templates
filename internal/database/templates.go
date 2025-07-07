package database

import (
	"database/sql"
	"fmt"
)

type Template struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	AuthorURL   string `json:"author_url"`
	CloneURL    string `json:"clone_url"`
	Description string `json:"description"`
}

type ListTemplatesArg struct {
	// The search phrase to use. Leave blank to list all.
	Search string
}

func ListTemplates(arg ListTemplatesArg) ([]Template, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}

	var (
		rows *sql.Rows
	)

	if len(arg.Search) > 0 {
		search := fmt.Sprintf("%%%s%%", arg.Search)
		rows, err = conn.Query(
			`
			 select * from templates
			 where name like ?
			 or clone_url like ?
			 or author like ?
			 or description like ?;
			`,
			search, search, search, search,
		)
	} else {
		rows, err = conn.Query(
			"select id, name, author, author_url, clone_url, description from templates",
		)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Template
	for rows.Next() {
		var t Template
		if err := rows.Scan(
			&t.ID, &t.Name, &t.Author, &t.AuthorURL, &t.CloneURL, &t.Description,
		); err != nil {
			return nil, err
		}
		results = append(results, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
