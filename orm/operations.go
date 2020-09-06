package orm

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// TruncateTable truncates a table
func TruncateTable(db *gorm.DB, model interface{}) error {
	return db.Debug().Exec(fmt.Sprintf("TRUNCATE TABLE %s", TableName(db, model))).Error
}

var dropTemplate = `IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].%s') AND type in (N'U')) DROP TABLE [dbo].%s;`

// DropTableMsSQL for create a DROP IF EXISTS statement in MsSQL
func DropTableMsSQL(db *gorm.DB, model interface{}) error {
	tableName := TableName(db, model)
	return db.Debug().Exec(fmt.Sprintf(dropTemplate, tableName, tableName)).Error
}

// Indexable interface for index
type Indexable interface {
	TableName() string
	UniqueColumns() []string
}

// ErrNotIndexable when a model is not indexable
var ErrNotIndexable = errors.New("model not indexable")

// CreateMsSQLIndex creates an index via mssql syntax
func CreateMsSQLIndex(db *gorm.DB, model interface{}) error {
	v, ok := model.(Indexable)
	if !ok {
		return ErrNotIndexable
	}

	stmt := fmt.Sprintf(
		"CREATE UNIQUE CLUSTERED INDEX %s_ix ON %s (%s)",
		v.TableName(),
		v.TableName(),
		strings.Join(v.UniqueColumns(), ","),
	)
	return db.LogMode(true).Exec(stmt).Error
}
