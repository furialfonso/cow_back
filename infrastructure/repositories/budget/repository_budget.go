package budget

import (
	"context"
	"fmt"

	"shared-wallet-service/infrastructure/database"

	"shared-wallet-service/domain/budget"
	"shared-wallet-service/domain/budget/dto"

	query "shared-wallet-service/infrastructure/database/queries/budget"
)

type budgetRepository struct {
	db database.IDataBase
}

func NewBudgetRepository(db database.IDataBase) budget.IBudgetRepository {
	return &budgetRepository{
		db: db,
	}
}

func (gr *budgetRepository) GetAll(ctx context.Context) ([]dto.Budget, error) {
	var budgets []dto.Budget
	rs, err := gr.db.GetRead().QueryContext(ctx, query.GetAll)
	if err != nil {
		return budgets, err
	}
	for rs.Next() {
		var budget dto.Budget
		if err := rs.Scan(
			&budget.ID,
			&budget.Code,
			&budget.Debt,
			&budget.CreatedAt,
		); err != nil {
			return budgets, err
		}
		budgets = append(budgets, budget)
	}
	return budgets, nil
}

func (gr *budgetRepository) GetByCode(ctx context.Context, code string) (dto.Budget, error) {
	var budgetModel dto.Budget
	rs, err := gr.db.GetRead().QueryContext(ctx, query.GetByCode, code)
	if err != nil {
		return budgetModel, err
	}
	exists := rs.Next()
	if !exists {
		return budgetModel, fmt.Errorf("budget %s not found", code)
	}
	if err := rs.Scan(
		&budgetModel.ID,
		&budgetModel.Code,
		&budgetModel.Debt,
		&budgetModel.CreatedAt,
	); err != nil {
		return budgetModel, err
	}
	return budgetModel, nil
}

func (gr *budgetRepository) Insert(ctx context.Context, code string) (int64, error) {
	rs, err := gr.db.GetWrite().ExecuteContext(ctx, query.Create, code)
	if err != nil {
		return 0, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (gr *budgetRepository) Delete(ctx context.Context, code string) error {
	_, err := gr.db.GetWrite().ExecuteContext(ctx, query.Delete, code)
	if err != nil {
		return err
	}
	return nil
}

func (gr *budgetRepository) UpdateDebtByCode(ctx context.Context, budgetModel dto.Budget) error {
	_, err := gr.db.GetWrite().ExecuteContext(ctx, query.Update, budgetModel.Debt, budgetModel.Code)
	if err != nil {
		return err
	}
	return nil
}
