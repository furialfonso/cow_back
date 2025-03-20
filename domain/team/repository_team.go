package team

import (
	"context"
	"shared-wallet-service/domain/budget/dto"
	dto2 "shared-wallet-service/domain/team/dto"
)

type ITeamRepository interface {
	GetTeamByBudget(ctx context.Context, code string) ([]string, error)
	GetTeamsByUser(ctx context.Context, code string) ([]dto.Budget, error)
	ExistUserInTeam(ctx context.Context, id string) (bool, error)
	ComposeTeam(ctx context.Context, team dto2.Team) (int64, error)
	DecomposeTeam(ctx context.Context, team dto2.Team) error
}
