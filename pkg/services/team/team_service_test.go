package team

import (
	"context"
	"cow_back/api/dto/response"
	"cow_back/mocks"
	"testing"

	"github.com/go-playground/assert"
)

type mockTeamService struct {
	groupRepository *mocks.IGroupRepository
	teamRepository  *mocks.ITeamRepository
	userCache       *mocks.ICache
}

type teamMocks struct {
	teamService func(f *mockTeamService)
}

func Test_GetUsersByGroup(t *testing.T) {
	tests := []struct {
		name   string
		code   string
		mocks  teamMocks
		outPut response.UsersByTeamResponse
		expErr error
	}{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockTeamService{
				groupRepository: &mocks.IGroupRepository{},
				teamRepository:  &mocks.ITeamRepository{},
				userCache:       &mocks.ICache{},
			}
			tc.mocks.teamService(m)
			service := NewTeamService(m.groupRepository, m.teamRepository, m.userCache)
			team, err := service.GetTeamByGroup(context.Background(), tc.code)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.outPut, team)
		})
	}
}
