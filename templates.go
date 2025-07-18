package templates

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrTemplateExists   = errors.New("template exists")
	ErrTemplateNotFound = errors.New("template not found")
	ErrOfficialTemplate = errors.New("cannot delete official template")
)

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
			 or source like ?;
			`,
			search, search,
		)
	} else {
		rows, err = c.Query(
			"select * from templates",
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
			&t.ID,
			&t.Name,
			&t.Source,
			&t.Type,
			&t.IsOfficial,
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
		insert into templates(name, source, type, is_official) 
		values (?, ?, ?, ?)
		on conflict(source)
		do update set
			name        = excluded.name;
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tmpl.Name, tmpl.Source, tmpl.Type, tmpl.IsOfficial)
	return err
}

type CreateTemplate struct {
	Name       string       `json:"name"`
	Source     string       `json:"source"`
	Type       TemplateType `json:"type"`
	IsOfficial bool         `json:"-"`
}

func (c *Connection) CreateTemplate(tmpl CreateTemplate) error {
	// Insert data
	stmt, err := c.Prepare(`
		insert into templates(name, source, type, is_official) 
		values (?, ?, ?, ?);
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tmpl.Name, tmpl.Source, tmpl.Type, tmpl.IsOfficial)
	return err
}

func (c *Connection) GetTemplateByName(name string) (*Template, error) {
	rows, err := c.Query(`
		select * from templates
		where name = ?
		limit 1
	`, name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []Template
	for rows.Next() {
		var t Template
		if err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Source,
			&t.Type,
			&t.IsOfficial,
		); err != nil {
			return nil, err
		}
		results = append(results, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(results) < 1 {
		return nil, ErrTemplateNotFound
	}

	template := results[0]
	return &template, nil
}

func (c *Connection) DeleteTemplate(name string) error {
	template, err := c.GetTemplateByName(name)
	if err != nil {
		return err
	}

	if template.IsOfficial {
		return ErrOfficialTemplate
	}

	_, err = c.Exec(`
		delete from templates
		where name = ?
	`, name)

	return err
}
