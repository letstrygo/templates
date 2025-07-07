package database

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/samber/lo"
)

const (
	CSVFile string = "./templates.csv"
)

func Migrate() error {
	// Clear any existing database.
	_ = os.Remove(DatabaseURL)

	// Open CSV file
	f, err := os.Open(CSVFile)
	if err != nil {
		return err
	}
	defer f.Close()

	// Parse CSV
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Skip header and map to Person structs
	templates := lo.Map(records[1:], func(row []string, _ int) Template {
		id, _ := strconv.Atoi(row[0])
		return Template{
			ID:          id,
			Name:        row[1],
			Author:      row[2],
			AuthorURL:   row[3],
			CloneURL:    row[4],
			Description: row[5],
		}
	})

	// Validate Templates
	err = validateTemplates(templates)
	if err != nil {
		return err
	}

	// Open SQLite database
	db, err := GetConnection()
	if err != nil {
		return err
	}

	// Create table
	_, err = db.Exec(`
        create table if not exists templates (
            id integer primary key,
            name text,
            author text,
            author_url text,
            clone_url text,
            description text
        );
    `)
	if err != nil {
		return err
	}

	// Insert data
	stmt, err := db.Prepare("insert into templates(id, name, author, author_url, clone_url, description) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, t := range templates {
		_, err := stmt.Exec(t.ID, t.Name, t.Author, t.AuthorURL, t.CloneURL, t.Description)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateTemplates(templates []Template) error {
	// Check for duplicate names
	names := lo.Map(templates, func(t Template, _ int) string { return t.Name })
	if len(lo.Uniq(names)) != len(names) {
		return errors.New("duplicate template name found")
	}

	// Check for duplicate clone URLs
	cloneURLs := lo.Map(templates, func(t Template, _ int) string { return t.CloneURL })
	if len(lo.Uniq(cloneURLs)) != len(cloneURLs) {
		return errors.New("duplicate clone_url found")
	}

	return nil
}
