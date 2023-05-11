package repository

import (
	"context"
	"docker-go-project/pkg/platform/database"
)

type repository struct {
	dataBase database.IDataBase
}

func NewRepository(dataBase database.IDataBase) IRepository {
	return &repository{
		dataBase: dataBase,
	}
}

func (r *repository) Get(ctx context.Context) (string, error) {
	rs, err := r.dataBase.GetRead().QueryContext(ctx, "select 'hi bb' text")
	if err != nil {
		return "", err
	}
	rs.Next()
	var value string
	rs.Scan(&value)
	return value, nil
}
