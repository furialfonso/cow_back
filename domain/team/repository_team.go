package team

import (
	"context"

	dtoBudget "shared-wallet-service/domain/budget/dto"
	dtoTeam "shared-wallet-service/domain/team/dto"
)

type ITeamRepository interface {
	GetTeamByBudget(ctx context.Context, code string) ([]string, error)
	GetTeamsByUser(ctx context.Context, code string) ([]dtoBudget.Budget, error)
	ExistUserInTeam(ctx context.Context, id string) (bool, error)
	ComposeTeam(ctx context.Context, team dtoTeam.Team) (int64, error)
	DecomposeTeam(ctx context.Context, team dtoTeam.Team) error
}
