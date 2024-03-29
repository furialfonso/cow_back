package database

import (
	"database/sql"

	sql2 "cow_back/pkg/platform/sql"

	_ "github.com/go-sql-driver/mysql"
)

type IDataBase interface {
	GetRead() sql2.IRead
	GetWrite() sql2.IWrite
}

type dataBase struct {
	conR   *sql.DB
	conW   *sql.DB
	nameDB string
}

func NewDataBase(nameDB string) IDataBase {
	return &dataBase{
		nameDB: nameDB,
		conR:   newConfigDB(nameDB, "R"),
		conW:   newConfigDB(nameDB, "W"),
	}
}

func (db *dataBase) GetRead() sql2.IRead {
	return sql2.NewSqlRead(db.conR)
}

func (db *dataBase) GetWrite() sql2.IWrite {
	return sql2.NewSqlWrite(db.conW)
}
