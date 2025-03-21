package team

import (
	"context"
	"fmt"

	"shared-wallet-service/domain/budget"
	"shared-wallet-service/domain/team"
	"shared-wallet-service/domain/team/dto"

	"shared-wallet-service/domain/user"

	"shared-wallet-service/interfaces/dto/request"
	"shared-wallet-service/interfaces/dto/response"

	"github.com/google/logger"
)

type ITeamUseCase interface {
	GetTeamByBudget(ctx context.Context, code string) (response.UsersByTeamResponse, error)
	GetTeamsByUser(ctx context.Context, userID string) (response.TeamsByUserResponse, error)
	ComposeTeam(ctx context.Context, code string, teamRequest request.TeamRequest) error
	DecomposeTeam(ctx context.Context, code string, teamRequest request.TeamRequest) error
}

type teamUseCase struct {
	budgetRepository budget.IBudgetRepository
	teamRepository   team.ITeamRepository
	cacheRepository  user.ICacheRepository
}

func NewTeamUseCase(budgetRepository budget.IBudgetRepository,
	teamRepository team.ITeamRepository,
	cacheRepository user.ICacheRepository,
) ITeamUseCase {
	return &teamUseCase{
		budgetRepository: budgetRepository,
		teamRepository:   teamRepository,
		cacheRepository:  cacheRepository,
	}
}

func (ts *teamUseCase) GetTeamByBudget(ctx context.Context, code string) (response.UsersByTeamResponse, error) {
	var usersByTeamResponse response.UsersByTeamResponse
	users, err := ts.teamRepository.GetTeamByBudget(ctx, code)
	if err != nil {
		return usersByTeamResponse, err
	}
	for _, userID := range users {
		user, exists := ts.cacheRepository.GetUser(userID)
		if !exists {
			continue
		}
		usersByTeamResponse.Users = append(usersByTeamResponse.Users, response.UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
			NickName: user.NickName,
		})
	}
	usersByTeamResponse.BudgetName = code

	return usersByTeamResponse, nil
}

func (ts *teamUseCase) GetTeamsByUser(ctx context.Context, userID string) (response.TeamsByUserResponse, error) {
	var teamsByUserResponse response.TeamsByUserResponse
	budgets, err := ts.teamRepository.GetTeamsByUser(ctx, userID)
	if err != nil {
		return teamsByUserResponse, err
	}
	for _, budget := range budgets {
		teamsByUserResponse.Teams = append(teamsByUserResponse.Teams, response.TeamResponse{
			Code: budget.Code,
			Debt: budget.Debt,
		})
	}
	return teamsByUserResponse, nil
}

func (ts *teamUseCase) ComposeTeam(ctx context.Context, code string, teamRequest request.TeamRequest) error {
	budget, err := ts.budgetRepository.GetByCode(ctx, code)
	if err != nil {
		logger.Error("error searching budget:", err)
		return err
	}

	for _, nickName := range teamRequest.Users {
		user, exists := ts.cacheRepository.GetUserByNickName(nickName)
		if !exists {
			logger.Infof("user %s not found", nickName)
			continue
		}
		exist, err := ts.teamRepository.ExistUserInTeam(ctx, user.ID)
		if err != nil {
			logger.Error("error searching user", err)
			return err
		}
		if exist {
			logger.Infof("user %s already exist in the team %s", nickName, code)
			continue
		}
		_, err = ts.teamRepository.ComposeTeam(ctx, dto.Team{
			BudgetID: budget.ID,
			UserID:   user.ID,
		})
		if err != nil {
			logger.Error(fmt.Sprintf("error adding user:%v ", nickName), err)
			return err
		}
	}
	return nil
}

func (ts *teamUseCase) DecomposeTeam(ctx context.Context, code string, teamRequest request.TeamRequest) error {
	budget, err := ts.budgetRepository.GetByCode(ctx, code)
	if err != nil {
		logger.Error("error searching budget:", err)
		return err
	}
	for _, nickName := range teamRequest.Users {
		user, exists := ts.cacheRepository.GetUserByNickName(nickName)
		if !exists {
			logger.Infof("user %s not found", nickName)
			continue
		}

		exist, err := ts.teamRepository.ExistUserInTeam(ctx, user.ID)
		if err != nil {
			logger.Error("error searching user", err)
			return err
		}
		if !exist {
			logger.Infof("user %s doesn't exist in the team %s", nickName, code)
			continue
		}
		err = ts.teamRepository.DecomposeTeam(ctx, dto.Team{
			BudgetID: budget.ID,
			UserID:   user.ID,
		})
		if err != nil {
			logger.Error(fmt.Sprintf("error deleting user:%v ", nickName), err)
			return err
		}
	}

	return nil
}
