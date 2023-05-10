package database

import (
	sql2 "docker-go-project/pkg/platform/sql"

	_ "github.com/go-sql-driver/mysql"
)

type IDataBase interface {
	GetRead() sql2.IRead
	GetWrite() sql2.IWrite
}

type dataBase struct {
	read  sql2.IRead
	write sql2.IWrite
}

func NewDataBase(nameDB string) IDataBase {
	return &dataBase{
		read:  sql2.NewSqlRead(newConfigDB(nameDB, "R")),
		write: sql2.NewSqlWrite(newConfigDB(nameDB, "W")),
	}
}

func (db *dataBase) GetRead() sql2.IRead {
	return db.read
}

func (db *dataBase) GetWrite() sql2.IWrite {
	return db.write
}
