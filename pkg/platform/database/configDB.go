package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
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
		user:     fmt.Sprintf("%s_%s", os.Getenv("USER_DB"), action),
		password: os.Getenv("PASSWORD_DB"),
		host:     os.Getenv("HOST_DB"),
		port:     os.Getenv("PORT_DB"),
		schema:   os.Getenv("SCHEMA_DB"),
	}
	stream := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.user, config.password, config.host, config.port, config.schema)
	pool, err := sql.Open(driver, stream)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = pool.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	return pool
}
