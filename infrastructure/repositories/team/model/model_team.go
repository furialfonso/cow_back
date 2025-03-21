package model

import "shared-wallet-service/domain/team/dto"

type Team struct {
	UserID    string `db:"user_id"`
	CreatedAt string `db:"created_at"`
	ID        int64  `db:"id"`
	BudgetID  int64  `db:"budget_id"`
}

func (model *Team) modelToDto() dto.Team {
	return dto.Team{
		UserID:    model.UserID,
		CreatedAt: model.CreatedAt,
		ID:        model.ID,
		BudgetID:  model.BudgetID,
	}
}
