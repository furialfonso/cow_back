package budget

import (
	"context"

	"shared-wallet-service/domain/budget"
	"shared-wallet-service/domain/budget/dto"

	"shared-wallet-service/interfaces/dto/request"
	"shared-wallet-service/interfaces/dto/response"
)

type IBudgetUseCase interface {
	GetAll(ctx context.Context) ([]response.BudgetResponse, error)
	GetByCode(ctx context.Context, code string) (response.BudgetResponse, error)
	Create(ctx context.Context, budgetRequest request.BudgetRequest) error
	Delete(ctx context.Context, code string) error
	UpdateDebtByCode(ctx context.Context, budgetRequest request.BudgetRequest) error
}

type budgetUseCase struct {
	budgetRepository budget.IBudgetRepository
}

func NewBudgetUseCase(budgetRepository budget.IBudgetRepository) IBudgetUseCase {
	return &budgetUseCase{
		budgetRepository: budgetRepository,
	}
}

func (gs *budgetUseCase) GetAll(ctx context.Context) ([]response.BudgetResponse, error) {
	budgets, err := gs.budgetRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return mapBudgetsToResponse(budgets), nil
}

func mapBudgetsToResponse(budgets []dto.Budget) []response.BudgetResponse {
	var responses []response.BudgetResponse
	for _, budget := range budgets {
		responses = append(responses, response.BudgetResponse{
			Code: budget.Code,
		})
	}
	return responses
}

func (gs *budgetUseCase) GetByCode(ctx context.Context, code string) (response.BudgetResponse, error) {
	var budgetResponse response.BudgetResponse
	rs, err := gs.budgetRepository.GetByCode(ctx, code)
	if err != nil {
		return budgetResponse, err
	}
	budgetResponse.Code = rs.Code
	return budgetResponse, nil
}

func (gs *budgetUseCase) Create(ctx context.Context, budgetRequest request.BudgetRequest) error {
	_, err := gs.budgetRepository.Insert(ctx, budgetRequest.Code)
	if err != nil {
		return err
	}
	return nil
}

func (gs *budgetUseCase) Delete(ctx context.Context, code string) error {
	err := gs.budgetRepository.Delete(ctx, code)
	if err != nil {
		return err
	}
	return nil
}

func (gs *budgetUseCase) UpdateDebtByCode(ctx context.Context, budgetRequest request.BudgetRequest) error {
	err := gs.budgetRepository.UpdateDebtByCode(ctx, dto.Budget{
		Code: budgetRequest.Code,
		Debt: budgetRequest.Debt,
	})
	if err != nil {
		return err
	}
	return nil
}
