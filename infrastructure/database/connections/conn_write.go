package connections

import (
	"database/sql"
	"sync"

	iwrite "shared-wallet-service/infrastructure/database/interfaces/write"
	"shared-wallet-service/infrastructure/database/write"

	_ "github.com/go-sql-driver/mysql"
)

var (
	writeConn *sql.DB
	onceWrite sync.Once
)

type writeDataBase struct {
	writeConn *sql.DB
	nameDB    string
}

func NewWriteDataBase(nameDB string) iwrite.IWriteDataBase {
	return &writeDataBase{
		nameDB:    nameDB,
		writeConn: getWriteConnection(nameDB),
	}
}

func getWriteConnection(nameDB string) *sql.DB {
	onceWrite.Do(func() {
		writeConn = NewDBConnection(nameDB, "W")
	})
	return writeConn
}

func (db *writeDataBase) Write() iwrite.IWrite {
	return write.NewSqlWrite(db.writeConn)
}
