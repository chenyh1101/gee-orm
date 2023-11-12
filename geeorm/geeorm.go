package geeorm

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/log"
	Session "geeorm/session"
	_ "github.com/mattn/go-sqlite3"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(dbType, dbPath string) (e *Engine, err error) {
	db, err := sql.Open(dbType, dbPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}
	dial, ok := dialect.GetDialect(dbType)
	if !ok {
		log.Errorf("dialect %s Not Found", dbType)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("connect database success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("failed to close database")
		return
	}
	log.Info("Close database success")
}

func (e *Engine) NewSession() *Session.Session {
	return Session.New(e.db, e.dialect)
}
