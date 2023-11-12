package schema

import (
	"fmt"
	"geeorm/dialect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.TableName != "User" || len(schema.FieldNames) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
	for _, v := range schema.Fields {
		fmt.Println(v.Name)
		fmt.Println(v.Type)
	}
	fmt.Println(schema.FieldNames)
}
