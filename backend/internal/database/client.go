package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Client struct {
	db *gorm.DB
}

func NewClient(config Config) (*Client, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Name, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Client{
		db: db,
	}, nil
}

func NewClientFromEnv() (*Client, error) {
	config, err := NewConfigFromEnv()
	if err != nil {
		return nil, err
	}

	return NewClient(config)
}

func (c *Client) SaveAreas(areas []Area) error {
	return c.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&areas).Error
}

func (c *Client) SaveCompetitions(competitions []Competition) error {
	return c.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&competitions).Error
}

func (c *Client) SaveEdition(edition Edition) error {
	return c.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&edition).Error
}

func (c *Client) SaveTeams(teams []Team) error {
	return c.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&teams).Error
}

func (c *Client) SaveMatches(matches []Match) error {
	return c.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&matches).Error
}

func (c *Client) SaveGoalScorers(goalScorers []GoalScorer) error {
	return c.db.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&goalScorers).Error
}

func (c *Client) DB() *gorm.DB {
	return c.db
}
