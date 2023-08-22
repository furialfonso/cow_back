package team

import (
	"context"
	"docker-go-project/api/dto/request"
	"docker-go-project/api/dto/response"
	"docker-go-project/pkg/repository/group"
	"docker-go-project/pkg/repository/team"
	"docker-go-project/pkg/repository/user"
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
	userRepository  user.IUserRepository
	teamRepository  team.ITeamRepository
}

func NewTeamService(groupRepository group.IGroupRepository,
	userRepository user.IUserRepository,
	teamRepository team.ITeamRepository) ITeamService {
	return &teamService{
		groupRepository: groupRepository,
		userRepository:  userRepository,
		teamRepository:  teamRepository,
	}
}

func (ts *teamService) GetUsersByGroup(ctx context.Context, code string) (response.TeamUsersResponse, error) {
	var teamUserResponse response.TeamUsersResponse
	users, err := ts.teamRepository.GetUsersByGroup(ctx, code)
	if err != nil {
		return teamUserResponse, err
	}
	for _, user := range users {
		us, err := ts.userRepository.GetByNickName(ctx, user)
		if err != nil {
			return teamUserResponse, err
		}
		teamUserResponse.Users = append(teamUserResponse.Users, response.UserResponse{
			Name:           us.Name,
			SecondName:     us.SecondName,
			LastName:       us.LastName,
			SecondLastName: us.SecondLastName,
			Email:          us.Email,
			NickName:       us.NickName,
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
		user, err := ts.userRepository.GetByNickName(ctx, nickName)
		if err != nil {
			logger.Error("error searching user", err)
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
		user, err := ts.userRepository.GetByNickName(ctx, nickName)
		if err != nil {
			logger.Error("error searching user", err)
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
