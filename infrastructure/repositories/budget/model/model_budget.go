package model

import "shared-wallet-service/domain/budget/dto"

type Budget struct {
	Code      string `db:"code"`
	CreatedAt string `db:"created_at"`
	ID        int64  `db:"id"`
	Debt      int    `db:"debt"`
}

func (model *Budget) ModelToDto() dto.Budget {
	return dto.Budget{
		ID:        model.ID,
		Code:      model.Code,
		Debt:      model.Debt,
		CreatedAt: model.CreatedAt,
	}
}
