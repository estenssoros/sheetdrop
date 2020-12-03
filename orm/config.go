package orm

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dialectPSQL = "psql"
)

// Config for database connection
type Config struct {
	Database string
	Host     string
	User     string
	Password string
	Dialect  string
}

func (c Config) String() string {
	ju, _ := json.MarshalIndent(c, "", " ")
	return string(ju)
}

// URL creates a connection url
func (c *Config) URL() string {
	switch c.Dialect {
	case dialectPSQL:
		return c.urlPSQL()
	default:
		return ""
	}
}

func (c *Config) urlPSQL() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s port=5432 sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Database,
	)
}

func (c *Config) dialector() (gorm.Dialector, error) {
	switch c.Dialect {
	case dialectPSQL:
		return postgres.Open(c.URL()), nil
	default:
		return nil, errors.Errorf("unknown dialect: %s", c.Dialect)
	}
}

func (c *Config) database() string {
	return c.Database
}
