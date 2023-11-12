package Session

import (
	"fmt"
	"geeorm/log"
	"geeorm/schema"
	"reflect"
	"strings"
)

func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.TableName) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("model is not set")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.refTable
	var columns []string
	for _, v := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", v.Name, v.Type, v.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.TableName, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().TableName)).Exec()
	return err
}
func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().TableName)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.refTable.TableName
}
