package Session

import (
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	err := s.DropTable()
	if err != nil {
		t.Fatal("failed to drop table user")
	}
	err = s.CreateTable()
	if err != nil {
		t.Fatal("failed to drop table user")
	}
	if !s.HasTable() {
		t.Fatal("failed to create table user")
	}
}
