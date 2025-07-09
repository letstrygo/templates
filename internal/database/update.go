package database

import "log"

func (c *Connection) Update() error {
	remoteDb, err := NewRemoteConnection()
	if err != nil {
		return err
	}

	tmpls, err := remoteDb.ListTemplates(ListTemplates{})
	if err != nil {
		return err
	}

	for _, t := range tmpls {
		err = c.UpsertTemplate(t)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}

	return remoteDb.Terminate()
}
