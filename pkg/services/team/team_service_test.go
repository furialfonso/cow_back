package team

import (
	"context"
	"docker-go-project/api/dto/response"
	"docker-go-project/mocks"
	"testing"

	"github.com/go-playground/assert"
)

type mockTeamService struct {
	groupRepository *mocks.IGroupRepository
	userRepository  *mocks.IUserRepository
	teamRepository  *mocks.ITeamRepository
}

type teamMocks struct {
	teamService func(f *mockTeamService)
}

func Test_GetUsersByGroup(t *testing.T) {
	tests := []struct {
		name   string
		code   string
		mocks  teamMocks
		outPut response.TeamUsersResponse
		expErr error
	}{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockTeamService{
				groupRepository: &mocks.IGroupRepository{},
				teamRepository:  &mocks.ITeamRepository{},
				userRepository:  &mocks.IUserRepository{},
			}
			tc.mocks.teamService(m)
			service := NewTeamService(m.groupRepository, m.userRepository, m.teamRepository)
			team, err := service.GetUsersByGroup(context.Background(), tc.code)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.outPut, team)
		})
	}
}
