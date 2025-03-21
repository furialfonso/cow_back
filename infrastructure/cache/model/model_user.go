package model

import "shared-wallet-service/domain/user/dto"

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	LastName string `db:"last_name"`
	Email    string `db:"email"`
	NickName string `db:"nick_name"`
}

func (model *User) ModelToDto() dto.User {
	return dto.User{
		ID:       model.ID,
		Name:     model.Name,
		LastName: model.LastName,
		Email:    model.Email,
		NickName: model.NickName,
	}
}
