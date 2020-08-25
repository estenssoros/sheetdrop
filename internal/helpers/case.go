package helpers

import "github.com/iancoleman/strcase"

func CamelCase(s string) string {
	return strcase.ToCamel(s)
}
