package budget

import (
	"context"
	"fmt"

	"shared-wallet-service/domain/budget"
	"shared-wallet-service/domain/budget/dto"
	iread "shared-wallet-service/infrastructure/database/interfaces/read"
	iwrite "shared-wallet-service/infrastructure/database/interfaces/write"
	"shared-wallet-service/infrastructure/database/queries"
	"shared-wallet-service/infrastructure/repositories/budget/model"
)

type budgetRepository struct {
	readDataBase  iread.IReadDataBase
	writeDataBase iwrite.IWriteDataBase
}

func NewBudgetRepository(readDataBase iread.IReadDataBase,
	writeDataBase iwrite.IWriteDataBase,
) budget.IBudgetRepository {
	return &budgetRepository{
		readDataBase:  readDataBase,
		writeDataBase: writeDataBase,
	}
}

func (gr *budgetRepository) GetAll(ctx context.Context) ([]dto.Budget, error) {
	var budgets []dto.Budget
	rs, err := gr.readDataBase.Read().QueryContext(ctx, queries.GetAll)
	if err != nil {
		return budgets, err
	}
	for rs.Next() {
		var budgetModel model.Budget
		if err := rs.Scan(
			&budgetModel.ID,
			&budgetModel.Code,
			&budgetModel.Debt,
			&budgetModel.CreatedAt,
		); err != nil {
			return budgets, err
		}
		budgets = append(budgets, budgetModel.ModelToDto())
	}
	return budgets, nil
}

func (gr *budgetRepository) GetByCode(ctx context.Context, code string) (dto.Budget, error) {
	var budgetDTO dto.Budget
	rs, err := gr.readDataBase.Read().QueryContext(ctx, queries.GetByCode, code)
	if err != nil {
		return budgetDTO, err
	}
	exists := rs.Next()
	if !exists {
		return budgetDTO, fmt.Errorf("budget %s not found", code)
	}

	var budgetModel model.Budget
	if err := rs.Scan(
		&budgetModel.ID,
		&budgetModel.Code,
		&budgetModel.Debt,
		&budgetModel.CreatedAt,
	); err != nil {
		return budgetDTO, err
	}

	return budgetModel.ModelToDto(), nil
}

func (gr *budgetRepository) Insert(ctx context.Context, code string) (int64, error) {
	rs, err := gr.writeDataBase.Write().ExecuteContext(ctx, queries.Create, code)
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
	_, err := gr.writeDataBase.Write().ExecuteContext(ctx, queries.Delete, code)
	if err != nil {
		return err
	}
	return nil
}

func (gr *budgetRepository) UpdateDebtByCode(ctx context.Context, budgetModel dto.Budget) error {
	_, err := gr.writeDataBase.Write().ExecuteContext(ctx, queries.Update, budgetModel.Debt, budgetModel.Code)
	if err != nil {
		return err
	}
	return nil
}
