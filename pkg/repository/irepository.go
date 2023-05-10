package repository

type IRepository interface {
	Get() (string, error)
}
