package budget

import (
	"context"
	"shared-wallet-service/domain/budget/dto"
)

type IBudgetRepository interface {
	GetAll(ctx context.Context) ([]dto.Budget, error)
	GetByCode(ctx context.Context, code string) (dto.Budget, error)
	Insert(ctx context.Context, code string) (int64, error)
	Delete(ctx context.Context, code string) error
	UpdateDebtByCode(ctx context.Context, budget dto.Budget) error
}
