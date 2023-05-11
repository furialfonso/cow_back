package repository

import "context"

type IRepository interface {
	Get(ctx context.Context) (string, error)
}
