package database

func (c *Connection) Migrate() error {
	// Create table
	_, err := c.Exec(`
        create table if not exists templates (
            id integer primary key,
            name text,
            author text,
            author_url text,
            clone_url text,
            description text
        );

		create unique index if not exists idx_templates_unique_name
		on templates (name);

		create unique index if not exists idx_templates_unique_clone_url
		on templates (clone_url);
    `)

	return err
}
