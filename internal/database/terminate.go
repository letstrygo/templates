package database

import "os"

// Terminate will both close the connection and delete it's underlying database file.
func (c *Connection) Terminate() error {
	if c.DB != nil {
		err := c.Close()
		if err != nil {
			return err
		}

		c.DB = nil
	}

	return os.Remove(c.file)
}
