package database

import (
	"database/sql"
	"docker-go-project/pkg/config"
	"fmt"
	"time"
)

const (
	driver = "mysql"
)

type configDB struct {
	user     string
	password string
	host     string
	port     string
	schema   string
}

func newConfigDB(nameDB string, action string) *sql.DB {
	config := configDB{
		user:     fmt.Sprintf("%s_%s", config.Get().UString(fmt.Sprintf("%s.username", nameDB)), action),
		password: config.Get().UString(fmt.Sprintf("%s.password", nameDB)),
		host:     config.Get().UString(fmt.Sprintf("%s.hostname", nameDB)),
		port:     config.Get().UString(fmt.Sprintf("%s.port", nameDB)),
		schema:   config.Get().UString(fmt.Sprintf("%s.schema", nameDB)),
	}
	stream := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.user, config.password, config.host, config.port, config.schema)
	db, err := sql.Open(driver, stream)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
