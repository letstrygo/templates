package database

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
		_ = c.UpsertTemplate(t)
	}

	return remoteDb.Terminate()
}
