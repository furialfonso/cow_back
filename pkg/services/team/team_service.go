package team

import (
	"context"
	"fmt"

	"cow_back/api/dto/request"
	"cow_back/api/dto/response"
	"cow_back/pkg/platform/cache"
	"cow_back/pkg/repository/group"
	"cow_back/pkg/repository/team"

	"github.com/google/logger"
)

type ITeamService interface {
	GetTeamByGroup(ctx context.Context, code string) (response.UsersByTeamResponse, error)
	GetTeamsByUser(ctx context.Context, userID string) (response.TeamsByUserResponse, error)
	ComposeTeam(ctx context.Context, code string, teamRequest request.TeamRequest) error
	DecomposeTeam(ctx context.Context, code string, teamRequest request.TeamRequest) error
}

type teamService struct {
	groupRepository group.IGroupRepository
	teamRepository  team.ITeamRepository
	userCache       cache.ICache
}

func NewTeamService(groupRepository group.IGroupRepository,
	teamRepository team.ITeamRepository,
	userCache cache.ICache,
) ITeamService {
	return &teamService{
		groupRepository: groupRepository,
		teamRepository:  teamRepository,
		userCache:       userCache,
	}
}

func (ts *teamService) GetTeamByGroup(ctx context.Context, code string) (response.UsersByTeamResponse, error) {
	var usersByTeamResponse response.UsersByTeamResponse
	users, err := ts.teamRepository.GetTeamByGroup(ctx, code)
	if err != nil {
		return usersByTeamResponse, err
	}
	for _, userID := range users {
		user, exists := ts.userCache.Get(userID)
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
	usersByTeamResponse.GroupName = code

	return usersByTeamResponse, nil
}

func (ts *teamService) GetTeamsByUser(ctx context.Context, userID string) (response.TeamsByUserResponse, error) {
	var teamsByUserResponse response.TeamsByUserResponse
	groups, err := ts.teamRepository.GetTeamsByUser(ctx, userID)
	if err != nil {
		return teamsByUserResponse, err
	}
	for _, group := range groups {
		teamsByUserResponse.Teams = append(teamsByUserResponse.Teams, response.TeamResponse{
			Code: group.Code,
			Debt: group.Debt,
		})
	}
	return teamsByUserResponse, nil
}

func (ts *teamService) ComposeTeam(ctx context.Context, code string, teamRequest request.TeamRequest) error {
	group, err := ts.groupRepository.GetByCode(ctx, code)
	if err != nil {
		logger.Error("error searching group:", err)
		return err
	}

	for _, nickName := range teamRequest.Users {
		user, exists := ts.userCache.GetByNickName(nickName)
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
		_, err = ts.teamRepository.ComposeTeam(ctx, team.Team{
			GroupID: group.ID,
			UserID:  user.ID,
		})
		if err != nil {
			logger.Error(fmt.Sprintf("error adding user:%v ", nickName), err)
			return err
		}
	}
	return nil
}

func (ts *teamService) DecomposeTeam(ctx context.Context, code string, teamRequest request.TeamRequest) error {
	group, err := ts.groupRepository.GetByCode(ctx, code)
	if err != nil {
		logger.Error("error searching group:", err)
		return err
	}
	for _, nickName := range teamRequest.Users {
		user, exists := ts.userCache.GetByNickName(nickName)
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
		err = ts.teamRepository.DecomposeTeam(ctx, team.Team{
			GroupID: group.ID,
			UserID:  user.ID,
		})
		if err != nil {
			logger.Error(fmt.Sprintf("error deleting user:%v ", nickName), err)
			return err
		}
	}

	return nil
}
