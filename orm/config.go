package orm

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type config struct {
	Database string
	Host     string
	User     string
	Password string `json:"-"`
	Dialect  string
}

func (c config) String() string {
	ju, _ := json.MarshalIndent(c, "", " ")
	return string(ju)
}

func (c *config) driverName() string {
	switch c.Dialect {
	case "microsoft_sql":
		return "mssql"
	default:
		return c.Dialect
	}
}

func (c *config) URL() (string, error) {
	switch c.Dialect {
	case "mysql":
		return c.urlMySQL(), nil
	default:
		return "", errors.New("no dialect")
	}
}

func (c *config) urlPostgres() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.User,
		c.Password,
		c.Database,
	)
}
func (c *config) urlMySQL() string {
	return fmt.Sprintf(
		"%s:%s@(%s)/%s?parseTime=true",
		c.User,
		c.Password,
		c.Host,
		c.Database,
	)
}
func (c *config) urlMsSQL() string {
	return fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		c.User,
		c.Password,
		c.Host,
		c.Database,
	)
}
