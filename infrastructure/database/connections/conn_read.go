package connections

import (
	"database/sql"
	"sync"

	iread "shared-wallet-service/infrastructure/database/interfaces/read"
	"shared-wallet-service/infrastructure/database/read"

	_ "github.com/go-sql-driver/mysql"
)

var (
	readConn *sql.DB
	onceRead sync.Once
)

type readDataBase struct {
	readConn *sql.DB
	nameDB   string
}

func NewReadDataBase(nameDB string) iread.IReadDataBase {
	return &readDataBase{
		nameDB:   nameDB,
		readConn: getReadConnection(nameDB),
	}
}

func getReadConnection(nameDB string) *sql.DB {
	onceRead.Do(func() {
		readConn = NewDBConnection(nameDB, "R")
	})
	return readConn
}

func (db *readDataBase) Read() iread.IRead {
	return read.NewSqlRead(db.readConn)
}
