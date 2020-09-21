package orm

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dialectMySQL = "mysql"
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

func (c *config) URL() (string, error) {
	switch c.Dialect {
	case dialectMySQL:
		return c.urlMySQL(), nil
	default:
		return "", errors.New("no dialect")
	}
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

func (c *config) dialector() gorm.Dialector {
	switch c.Dialect {
	case dialectMySQL:
		return mysql.Open(c.urlMySQL())
	default:
		return nil
	}
}
