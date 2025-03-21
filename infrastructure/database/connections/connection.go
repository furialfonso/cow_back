package connections

import (
	"context"
	"database/sql"
	"time"
)

const (
	mysqlDriver = "mysql"
)

type configDB struct {
	user     string
	password string
	host     string
	port     string
	schema   string
}

func NewDBConnection(nameDB string, action string) *sql.DB {
	cfg := GetDBConfig(nameDB, action)
	conn, err := sql.Open(mysqlDriver, cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = conn.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	return conn
}
