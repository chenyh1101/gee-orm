package dialect

import (
	"reflect"
)

var dialectMap = map[string]Dialect{}

type Dialect interface {
	DataValueOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

func GetDialect(tableName string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[tableName]
	return
}
