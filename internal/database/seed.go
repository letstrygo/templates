package database

import (
	"encoding/csv"
	"errors"
	"os"
)

var (
	ErrIncompleteRow error = errors.New("incomplete row")
)

const (
	CSVFile string = "repository.csv"
)

func (c *Connection) Seed() error {
	f, err := os.Open(CSVFile)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	if len(records) < 2 {
		// No data
		return nil
	}

	var templates []CreateTemplate

	for _, row := range records[1:] {
		if len(row) < 5 {
			return ErrIncompleteRow
		}

		t := CreateTemplate{
			Name:        row[0],
			Author:      row[1],
			AuthorURL:   row[2],
			CloneURL:    row[3],
			Description: row[4],
			IsOfficial:  true,
		}
		templates = append(templates, t)
	}

	_, err = c.Exec("delete from templates;")
	if err != nil {
		return err
	}

	for _, t := range templates {
		err := c.CreateTemplate(t)
		if err != nil {
			return err
		}
	}

	return nil
}
