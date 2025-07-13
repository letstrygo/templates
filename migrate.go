package templates

func (c *Connection) Migrate() error {
	// Create table
	_, err := c.Exec(`
        create table if not exists templates (
            id          integer primary key,
            name        text,
            source      text,
            type        text,
            is_official boolean
        );

        create unique index if not exists idx_templates_unique_name
        on templates (name);

        create unique index if not exists idx_templates_unique_source
        on templates (source);
    `)

	return err
}
