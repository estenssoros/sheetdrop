package orm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
	"gorm.io/gorm"
)

var (
	writeLimit = 300
	argsLimit  = 2100
)

// IsSlice is you is a slice?
func IsSlice(m interface{}) bool {
	v := reflect.Indirect(reflect.ValueOf(m))
	return v.Kind() == reflect.Slice || v.Kind() == reflect.Array
}

func makeWildCardTuple(length int) string {
	cards := make([]string, length)
	for i := 0; i < length; i++ {
		cards[i] = "?"
	}
	return fmt.Sprintf("(%s)", strings.Join(cards, ","))
}

// InsertStatement create an insert statement
func InsertStatement(db *gorm.DB, models interface{}) string {
	scope := db.NewScope(models)
	fields := []string{}
	switch scope.Dialect().GetName() {
	case "mysql":
		for _, field := range scope.Fields() {
			fields = append(fields, fmt.Sprintf("`%s`", field.DBName))
		}
	default:
		for _, field := range scope.Fields() {
			fields = append(fields, field.DBName)
		}
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES ", scope.QuotedTableName(), strings.Join(fields, ","))
}

func interfaces(v interface{}) []interface{} {
	val := reflect.ValueOf(v).Elem()
	out := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		out[i] = val.Field(i).Interface()
	}
	return out
}

// TableName finds the models. tablename
func TableName(db *gorm.DB, model interface{}) string {
	return db.NewScope(model).QuotedTableName()
}

// CreateMany insert many
func CreateMany(db *gorm.DB, models interface{}) error {
	if !IsSlice(models) {
		return errors.New("must be slice")
	}
	wildCardTuple := makeWildCardTuple(len(db.NewScope(models).Fields()))
	insertStmt := InsertStatement(db, models)
	wildCards, valueArgs := []string{}, []interface{}{}
	v := reflect.Indirect(reflect.ValueOf(models))
	if v.Len() == 0 {
		return errors.New("no data")
	}
	dataLength := len(interfaces(v.Index(0).Interface()))
	bar := progressbar.Default(int64(v.Len()), TableName(db, models))
	var exec = func(db *gorm.DB, wildCards []string, valueArgs []interface{}) error {
		if err := db.LogMode(false).Exec(insertStmt+strings.Join(wildCards, ","), valueArgs...).Error; err != nil {
			db.Debug().Exec(insertStmt+strings.Join(wildCards, ","), valueArgs...)
			return err
		}
		return nil
	}
	for i := 0; i < v.Len(); i++ {
		bar.Add(1)
		wildCards = append(wildCards, wildCardTuple)
		valueArgs = append(valueArgs, interfaces(v.Index(i).Interface())...)
		if len(wildCards) == writeLimit || len(valueArgs)+dataLength >= argsLimit {
			if err := exec(db, wildCards, valueArgs); err != nil {
				return err
			}
			wildCards, valueArgs = nil, nil
		}
	}
	bar.Finish()
	if len(wildCards) > 0 {
		return exec(db, wildCards, valueArgs)
	}
	return nil
}
