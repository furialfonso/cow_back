package team

import (
	"context"

	"shared-wallet-service/infrastructure/database"
	query "shared-wallet-service/infrastructure/database/queries/team"

	"shared-wallet-service/domain/budget/dto"
	dto2 "shared-wallet-service/domain/team/dto"

	"shared-wallet-service/domain/team"
)

type teamRepository struct {
	db database.IDataBase
}

func NewTeamRepository(db database.IDataBase) team.ITeamRepository {
	return &teamRepository{
		db: db,
	}
}

func (tr *teamRepository) GetTeamByBudget(ctx context.Context, code string) ([]string, error) {
	rs, err := tr.db.GetRead().QueryContext(ctx, query.GetTeamByBudget, code)
	if err != nil {
		return nil, err
	}
	var users []string
	for rs.Next() {
		var user string
		if err := rs.Scan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (tr *teamRepository) GetTeamsByUser(ctx context.Context, code string) ([]dto.Budget, error) {
	rs, err := tr.db.GetRead().QueryContext(ctx, query.GetTeamsByUser, code)
	if err != nil {
		return nil, err
	}
	var budgets []dto.Budget
	for rs.Next() {
		var budget dto.Budget
		if err := rs.Scan(
			&budget.ID,
			&budget.Code,
			&budget.Debt,
			&budget.CreatedAt,
		); err != nil {
			return nil, err
		}
		budgets = append(budgets, budget)
	}
	return budgets, nil
}

func (tr *teamRepository) ExistUserInTeam(ctx context.Context, id string) (bool, error) {
	rs, err := tr.db.GetRead().QueryContext(ctx, query.GetUserByID, id)
	if err != nil {
		return false, err
	}
	rs.Next()
	var exist bool
	if err := rs.Scan(&exist); err != nil {
		return false, err
	}
	return exist, nil
}

func (tr *teamRepository) ComposeTeam(ctx context.Context, teamModel dto2.Team) (int64, error) {
	rs, err := tr.db.GetWrite().ExecuteContext(ctx, query.ComposeTeam, teamModel.BudgetID, teamModel.UserID)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}

func (tr *teamRepository) DecomposeTeam(ctx context.Context, teamModel dto2.Team) error {
	_, err := tr.db.GetWrite().ExecuteContext(ctx, query.DecomposeTeam, teamModel.BudgetID, teamModel.UserID)
	if err != nil {
		return err
	}
	return nil
}
