package main

import (
	"geeorm"
	"geeorm/log"
)

func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "/home/cyh/gee.db")
	defer engine.Close()
	session := engine.NewSession()
	_, _ = session.Raw("DROP TABLE IF EXISTS User1;").Exec()
	_, _ = session.Raw("CREATE TABLE IF NOT EXISTS User1(Name text);").Exec()
	_, _ = session.Raw("CREATE TABLE User1(Name text);").Exec()
	result, _ := session.Raw("INSERT INTO User1(`Name`) VALUES (?),(?)", "Tom", "Jack").Exec()
	count, _ := result.RowsAffected()
	log.Infof("Exec success,%d affected\n", count)

}
