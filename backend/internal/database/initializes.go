package database

import "slices"

var tables = []any{
	&SeasonPlayer{},
	&Player{},
	&Match{},
	&Team{},
	&Edition{},
	&Competition{},
	&Area{},
}

func (c *Client) DeleteDatabase() error {
	for _, table := range tables {
		if err := c.DB().Migrator().DropTable(table); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) CreateDatabase() error {
	for _, table := range slices.Backward(tables) {
		if err := c.DB().AutoMigrate(table); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) ResetDatabase() error {
	if err := c.DeleteDatabase(); err != nil {
		return err
	}

	return c.CreateDatabase()
}
