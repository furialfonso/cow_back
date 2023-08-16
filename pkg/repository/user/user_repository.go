package user

import (
	"context"
	"docker-go-project/pkg/platform/database"
	template "docker-go-project/pkg/platform/templates/user"
	"fmt"
)

type IUserRepository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetByNickName(ctx context.Context, nickName string) (User, error)
	Create(ctx context.Context, user User) (int64, error)
	Delete(ctx context.Context, code string) error
}

type userRepository struct {
	db database.IDataBase
}

func NewUserRepository(db database.IDataBase) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetAll(ctx context.Context) ([]User, error) {
	var users []User
	rs, err := ur.db.GetRead().QueryContext(ctx, template.GetAll)
	if err != nil {
		return users, err
	}
	for rs.Next() {
		var user User
		if err := rs.Scan(
			&user.ID,
			&user.Name,
			&user.SecondName,
			&user.LastName,
			&user.SecondLastName,
			&user.Email,
			&user.NickName,
			&user.CreatedAt,
		); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *userRepository) GetByNickName(ctx context.Context, nickName string) (User, error) {
	var user User
	rs, err := ur.db.GetRead().QueryContext(ctx, template.GetByCode, nickName)
	if err != nil {
		return user, err
	}
	exists := rs.Next()
	if !exists {
		return user, fmt.Errorf("user %s not found", nickName)
	}
	if err := rs.Scan(
		&user.ID,
		&user.Name,
		&user.SecondName,
		&user.LastName,
		&user.SecondLastName,
		&user.Email,
		&user.NickName,
		&user.CreatedAt,
	); err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepository) Create(ctx context.Context, user User) (int64, error) {
	rs, err := ur.db.GetWrite().ExecuteContext(ctx, template.Create, user.Name, user.SecondName, user.LastName, user.SecondLastName, user.Email, user.NickName)
	if err != nil {
		return 0, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ur *userRepository) Delete(ctx context.Context, nickName string) error {
	_, err := ur.db.GetWrite().ExecuteContext(ctx, template.Delete, nickName)
	if err != nil {
		return err
	}
	return nil
}
