package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrTemplateExists = errors.New("template exists")
)

type Template struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	AuthorURL   string `json:"author_url"`
	CloneURL    string `json:"clone_url"`
	Description string `json:"description"`
}

type ListTemplates struct {
	// The search phrase to use. Leave blank to list all.
	Search string
}

func (c *Connection) ListTemplates(arg ListTemplates) ([]Template, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if len(arg.Search) > 0 {
		search := fmt.Sprintf("%%%s%%", arg.Search)
		rows, err = c.Query(
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
		rows, err = c.Query(
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

func (c *Connection) UpsertTemplate(tmpl Template) error {
	// Insert data
	stmt, err := c.Prepare(`
		insert into templates(name, author, author_url, clone_url, description) 
		values (?, ?, ?, ?, ?)
		on conflict(clone_url)
		do update set
			name        = excluded.name,
			author      = excluded.author,
			author_url  = excluded.author_url,
			clone_url   = excluded.clone_url,
			description = excluded.description;
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tmpl.Name, tmpl.Author, tmpl.AuthorURL, tmpl.CloneURL, tmpl.Description)
	return err
}

type CreateTemplate struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	AuthorURL   string `json:"author_url"`
	CloneURL    string `json:"clone_url"`
	Description string `json:"description"`
}

func (c *Connection) CreateTemplate(tmpl CreateTemplate) error {
	// Insert data
	stmt, err := c.Prepare(`
		insert into templates(name, author, author_url, clone_url, description) 
		values (?, ?, ?, ?, ?);
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tmpl.Name, tmpl.Author, tmpl.AuthorURL, tmpl.CloneURL, tmpl.Description)
	return err
}
