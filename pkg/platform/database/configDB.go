package database

import (
	"context"
	"database/sql"
	"docker-go-project/pkg/config"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
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
	cfg := mysql.Config{
		User:   fmt.Sprintf("%s_%s", config.Get().UString(fmt.Sprintf("%s.user", nameDB)), action),
		Passwd: config.Get().UString(fmt.Sprintf("%s.password", nameDB)),
		Net:    "tcp",
		Addr: fmt.Sprintf("%s:%s", config.Get().UString(fmt.Sprintf("%s.host", nameDB)),
			config.Get().UString(fmt.Sprintf("%s.port", nameDB))),
		DBName:               config.Get().UString(fmt.Sprintf("%s.schema", nameDB)),
		AllowNativePasswords: true,
	}
	pool, err := sql.Open(driver, cfg.FormatDSN())

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
