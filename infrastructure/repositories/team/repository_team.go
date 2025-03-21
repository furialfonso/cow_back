package team

import (
	"context"

	iread "shared-wallet-service/infrastructure/database/interfaces/read"
	iwrite "shared-wallet-service/infrastructure/database/interfaces/write"
	"shared-wallet-service/infrastructure/database/queries"
	"shared-wallet-service/infrastructure/repositories/budget/model"

	dtoBudget "shared-wallet-service/domain/budget/dto"
	dtoUser "shared-wallet-service/domain/team/dto"

	"shared-wallet-service/domain/team"
)

type teamRepository struct {
	readDataBase  iread.IReadDataBase
	writeDataBase iwrite.IWriteDataBase
}

func NewTeamRepository(readDataBase iread.IReadDataBase,
	writeDataBase iwrite.IWriteDataBase,
) team.ITeamRepository {
	return &teamRepository{
		readDataBase:  readDataBase,
		writeDataBase: writeDataBase,
	}
}

func (tr *teamRepository) GetTeamByBudget(ctx context.Context, code string) ([]string, error) {
	rs, err := tr.readDataBase.Read().QueryContext(ctx, queries.GetTeamByBudget, code)
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

func (tr *teamRepository) GetTeamsByUser(ctx context.Context, code string) ([]dtoBudget.Budget, error) {
	rs, err := tr.readDataBase.Read().QueryContext(ctx, queries.GetTeamsByUser, code)
	if err != nil {
		return nil, err
	}
	var budgets []dtoBudget.Budget
	for rs.Next() {
		var budgetModel model.Budget
		if err := rs.Scan(
			&budgetModel.ID,
			&budgetModel.Code,
			&budgetModel.Debt,
			&budgetModel.CreatedAt,
		); err != nil {
			return nil, err
		}
		budgets = append(budgets, budgetModel.ModelToDto())
	}
	return budgets, nil
}

func (tr *teamRepository) ExistUserInTeam(ctx context.Context, id string) (bool, error) {
	rs, err := tr.readDataBase.Read().QueryContext(ctx, queries.GetUserByID, id)
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

func (tr *teamRepository) ComposeTeam(ctx context.Context, teamModel dtoUser.Team) (int64, error) {
	rs, err := tr.writeDataBase.Write().ExecuteContext(ctx, queries.ComposeTeam, teamModel.BudgetID, teamModel.UserID)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}

func (tr *teamRepository) DecomposeTeam(ctx context.Context, teamModel dtoUser.Team) error {
	_, err := tr.writeDataBase.Write().ExecuteContext(ctx, queries.DecomposeTeam, teamModel.BudgetID, teamModel.UserID)
	if err != nil {
		return err
	}
	return nil
}
