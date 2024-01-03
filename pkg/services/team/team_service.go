package team

import (
	"context"
	"docker-go-project/api/dto/request"
	"docker-go-project/api/dto/response"
	"docker-go-project/pkg/platform/cache"
	"docker-go-project/pkg/repository/group"
	"docker-go-project/pkg/repository/team"
	"fmt"

	"github.com/google/logger"
)

type ITeamService interface {
	GetUsersByGroup(ctx context.Context, code string) (response.TeamUsersResponse, error)
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
	userCache cache.ICache) ITeamService {
	return &teamService{
		groupRepository: groupRepository,
		teamRepository:  teamRepository,
		userCache:       userCache,
	}
}

func (ts *teamService) GetUsersByGroup(ctx context.Context, code string) (response.TeamUsersResponse, error) {
	var teamUserResponse response.TeamUsersResponse
	users, err := ts.teamRepository.GetUsersByGroup(ctx, code)
	if err != nil {
		return teamUserResponse, err
	}
	for _, userID := range users {
		user, exists := ts.userCache.Get(userID)
		if !exists {
			continue
		}
		teamUserResponse.Users = append(teamUserResponse.Users, response.UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
			NickName: user.NickName,
		})
	}
	teamUserResponse.GroupName = code

	return teamUserResponse, nil
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
