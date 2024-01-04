package team

import (
	"context"
	"errors"
	"testing"

	"cow_back/api/dto/request"
	"cow_back/api/dto/response"
	"cow_back/mocks"
	"cow_back/pkg/platform/cache"
	"cow_back/pkg/repository/group"
	"cow_back/pkg/repository/team"

	"github.com/go-playground/assert"
	"github.com/stretchr/testify/mock"
)

type mockTeamService struct {
	groupRepository *mocks.IGroupRepository
	teamRepository  *mocks.ITeamRepository
	userCache       *mocks.ICache
}

type teamMocks struct {
	teamService func(f *mockTeamService)
}

func Test_GetTeamByGroup(t *testing.T) {
	tests := []struct {
		expErr error
		mocks  teamMocks
		name   string
		code   string
		outPut response.UsersByTeamResponse
	}{
		{
			name: "doesn't exist second user",
			code: "test",
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.teamRepository.Mock.On("GetTeamByGroup", mock.Anything, "test").Return([]string{}, errors.New("some error"))
				},
			},
			expErr: errors.New("some error"),
		},
		{
			name: "doesn't exist second user",
			code: "test",
			mocks: teamMocks{
				func(f *mockTeamService) {
					users := []string{"1", "2"}
					f.teamRepository.Mock.On("GetTeamByGroup", mock.Anything, "test").Return(users, nil)
					for _, user := range users {
						if user == "1" {
							f.userCache.Mock.On("Get", user).Return(cache.User{
								ID:       "1",
								Name:     "diego",
								LastName: "fernandez",
								Email:    "diego@gmail.com",
								NickName: "diegof",
							}, true)
						} else {
							f.userCache.Mock.On("Get", user).Return(cache.User{}, false)
						}
					}
				},
			},
			outPut: response.UsersByTeamResponse{
				GroupName: "test",
				Users: []response.UserResponse{
					{
						ID:       "1",
						Name:     "diego",
						LastName: "fernandez",
						Email:    "diego@gmail.com",
						NickName: "diegof",
					},
				},
			},
		},
		{
			name: "full flow",
			code: "test",
			mocks: teamMocks{
				func(f *mockTeamService) {
					users := []string{"1", "2"}
					f.teamRepository.Mock.On("GetTeamByGroup", mock.Anything, "test").Return(users, nil)
					for _, user := range users {
						if user == "1" {
							f.userCache.Mock.On("Get", user).Return(cache.User{
								ID:       "1",
								Name:     "diego",
								LastName: "fernandez",
								Email:    "diego@gmail.com",
								NickName: "diegof",
							}, true)
						} else {
							f.userCache.Mock.On("Get", user).Return(cache.User{
								ID:       "2",
								Name:     "petunia",
								LastName: "avila",
								Email:    "petinia@gmail.com",
								NickName: "petuniaa",
							}, true)
						}
					}
				},
			},
			outPut: response.UsersByTeamResponse{
				GroupName: "test",
				Users: []response.UserResponse{
					{
						ID:       "1",
						Name:     "diego",
						LastName: "fernandez",
						Email:    "diego@gmail.com",
						NickName: "diegof",
					},
					{
						ID:       "2",
						Name:     "petunia",
						LastName: "avila",
						Email:    "petinia@gmail.com",
						NickName: "petuniaa",
					},
				},
			},
		},
	}
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

func Test_GetUsersByGroup(t *testing.T) {
	tests := []struct {
		expErr error
		mocks  teamMocks
		name   string
		userID string
		outPut response.TeamsByUserResponse
	}{
		{
			name:   "error GetTeamsByUser",
			userID: "1",
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.teamRepository.Mock.On("GetTeamsByUser", mock.Anything, "1").Return([]group.Group{}, errors.New("some error"))
				},
			},
			expErr: errors.New("some error"),
		},
		{
			name:   "without teams",
			userID: "1",
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.teamRepository.Mock.On("GetTeamsByUser", mock.Anything, "1").Return([]group.Group{}, nil)
				},
			},
			outPut: response.TeamsByUserResponse{},
		},
		{
			name:   "full flow",
			userID: "1",
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.teamRepository.Mock.On("GetTeamsByUser", mock.Anything, "1").Return([]group.Group{
						{
							Code:      "test",
							ID:        1,
							CreatedAt: "2021-01-01",
							Debt:      100,
						},
					}, nil)
				},
			},
			outPut: response.TeamsByUserResponse{
				Teams: []response.TeamResponse{
					{
						Code: "test",
						Debt: 100,
					},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockTeamService{
				groupRepository: &mocks.IGroupRepository{},
				teamRepository:  &mocks.ITeamRepository{},
				userCache:       &mocks.ICache{},
			}
			tc.mocks.teamService(m)
			service := NewTeamService(m.groupRepository, m.teamRepository, m.userCache)
			team, err := service.GetTeamsByUser(context.Background(), tc.userID)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.outPut, team)
		})
	}
}

func Test_ComposeTeam(t *testing.T) {
	tests := []struct {
		expErr      error
		mocks       teamMocks
		name        string
		code        string
		teamRequest request.TeamRequest
	}{
		{
			name: "error GetByCode",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{}, errors.New("some error"))
				},
			},
			expErr: errors.New("some error"),
		},
		{
			name: "user not found",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{
						Code:      "test",
						CreatedAt: "2021-01-01",
						ID:        1,
						Debt:      0,
					}, nil)
					users := []string{"diegof"}
					for _, user := range users {
						f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
							ID:       "A",
							Name:     "diego",
							LastName: "fernandez",
							Email:    "diego@gmail.com",
							NickName: "diegof",
						}, false)
					}
				},
			},
		},
		{
			name: "error ExistUserInTeam",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{
						Code:      "test",
						CreatedAt: "2021-01-01",
						ID:        1,
						Debt:      0,
					}, nil)
					users := []string{"diegof"}
					for _, user := range users {
						f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
							ID:       "A",
							Name:     "diego",
							LastName: "fernandez",
							Email:    "diego@gmail.com",
							NickName: "diegof",
						}, true)
						f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "A").Return(false, errors.New("some error"))
					}
				},
			},
			expErr: errors.New("some error"),
		},
		{
			name: "error ComposeTeam",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{
						Code:      "test",
						CreatedAt: "2021-01-01",
						ID:        1,
						Debt:      0,
					}, nil)
					users := []string{"diegof"}
					for _, user := range users {
						f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
							ID:       "A",
							Name:     "diego",
							LastName: "fernandez",
							Email:    "diego@gmail.com",
							NickName: "diegof",
						}, true)
						f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "A").Return(false, nil)
						f.teamRepository.Mock.On("ComposeTeam", mock.Anything, team.Team{
							GroupID: 1,
							UserID:  "A",
						}).Return(int64(1), errors.New("some error"))
					}
				},
			},
			expErr: errors.New("some error"),
		},
		{
			name: "user already exist in the team",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof", "petuniaa"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{
						Code:      "test",
						CreatedAt: "2021-01-01",
						ID:        1,
						Debt:      0,
					}, nil)
					users := []string{"diegof", "petuniaa"}
					for _, user := range users {
						if user == "diegof" {
							f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
								ID:       "A",
								Name:     "diego",
								LastName: "fernandez",
								Email:    "diego@gmail.com",
								NickName: "diegof",
							}, true)
							f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "A").Return(false, nil)
							f.teamRepository.Mock.On("ComposeTeam", mock.Anything, team.Team{
								GroupID: 1,
								UserID:  "A",
							}).Return(int64(1), nil)
						} else {
							f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
								ID:       "B",
								Name:     "petunia",
								LastName: "avila",
								Email:    "petunia@gmail.com",
								NickName: "petuniaa",
							}, true)
							f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "B").Return(true, nil)
						}
					}
				},
			},
		},
		{
			name: "full flow",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof", "petuniaa"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{
						Code:      "test",
						CreatedAt: "2021-01-01",
						ID:        1,
						Debt:      0,
					}, nil)
					users := []string{"diegof", "petuniaa"}
					for _, user := range users {
						if user == "diegof" {
							f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
								ID:       "A",
								Name:     "diego",
								LastName: "fernandez",
								Email:    "diego@gmail.com",
								NickName: "diegof",
							}, true)
							f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "A").Return(false, nil)
							f.teamRepository.Mock.On("ComposeTeam", mock.Anything, team.Team{
								GroupID: 1,
								UserID:  "A",
							}).Return(int64(1), nil)
						} else {
							f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
								ID:       "B",
								Name:     "petunia",
								LastName: "avila",
								Email:    "petunia@gmail.com",
								NickName: "petuniaa",
							}, true)
							f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "B").Return(false, nil)
							f.teamRepository.Mock.On("ComposeTeam", mock.Anything, team.Team{
								GroupID: 1,
								UserID:  "B",
							}).Return(int64(1), nil)
						}
					}
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockTeamService{
				groupRepository: &mocks.IGroupRepository{},
				teamRepository:  &mocks.ITeamRepository{},
				userCache:       &mocks.ICache{},
			}
			tc.mocks.teamService(m)
			service := NewTeamService(m.groupRepository, m.teamRepository, m.userCache)
			err := service.ComposeTeam(context.Background(), tc.code, tc.teamRequest)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
		})
	}
}

func Test_DecomposeTeam(t *testing.T) {
	tests := []struct {
		expErr      error
		mocks       teamMocks
		name        string
		code        string
		teamRequest request.TeamRequest
	}{
		{
			name: "error GetByCode",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{}, errors.New("some error"))
				},
			},
			expErr: errors.New("some error"),
		},
		{
			name: "user not found",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{
						Code:      "test",
						CreatedAt: "2021-01-01",
						ID:        1,
						Debt:      0,
					}, nil)
					users := []string{"diegof"}
					for _, user := range users {
						f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
							ID:       "A",
							Name:     "diego",
							LastName: "fernandez",
							Email:    "diego@gmail.com",
							NickName: "diegof",
						}, false)
					}
				},
			},
		},
		{
			name: "error ExistUserInTeam",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{
						Code:      "test",
						CreatedAt: "2021-01-01",
						ID:        1,
						Debt:      0,
					}, nil)
					users := []string{"diegof"}
					for _, user := range users {
						f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
							ID:       "A",
							Name:     "diego",
							LastName: "fernandez",
							Email:    "diego@gmail.com",
							NickName: "diegof",
						}, true)
						f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "A").Return(false, errors.New("some error"))
					}
				},
			},
			expErr: errors.New("some error"),
		},
		{
			name: "error ComposeTeam",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{
						Code:      "test",
						CreatedAt: "2021-01-01",
						ID:        1,
						Debt:      0,
					}, nil)
					users := []string{"diegof"}
					for _, user := range users {
						f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
							ID:       "A",
							Name:     "diego",
							LastName: "fernandez",
							Email:    "diego@gmail.com",
							NickName: "diegof",
						}, true)
						f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "A").Return(false, nil)
						f.teamRepository.Mock.On("DecomposeTeam", mock.Anything, team.Team{
							GroupID: 1,
							UserID:  "A",
						}).Return(errors.New("some error"))
					}
				},
			},
			expErr: errors.New("some error"),
		},
		{
			name: "user doesn't exist in the team",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof", "petuniaa"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{
						Code:      "test",
						CreatedAt: "2021-01-01",
						ID:        1,
						Debt:      0,
					}, nil)
					users := []string{"diegof", "petuniaa"}
					for _, user := range users {
						if user == "diegof" {
							f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
								ID:       "A",
								Name:     "diego",
								LastName: "fernandez",
								Email:    "diego@gmail.com",
								NickName: "diegof",
							}, true)
							f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "A").Return(true, nil)
							f.teamRepository.Mock.On("DecomposeTeam", mock.Anything, team.Team{
								GroupID: 1,
								UserID:  "A",
							}).Return(nil)
						} else {
							f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
								ID:       "B",
								Name:     "petunia",
								LastName: "avila",
								Email:    "petunia@gmail.com",
								NickName: "petuniaa",
							}, true)
							f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "B").Return(false, nil)
						}
					}
				},
			},
		},
		{
			name: "full flow",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diegof", "petuniaa"},
			},
			mocks: teamMocks{
				func(f *mockTeamService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "test").Return(group.Group{
						Code:      "test",
						CreatedAt: "2021-01-01",
						ID:        1,
						Debt:      0,
					}, nil)
					users := []string{"diegof", "petuniaa"}
					for _, user := range users {
						if user == "diegof" {
							f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
								ID:       "A",
								Name:     "diego",
								LastName: "fernandez",
								Email:    "diego@gmail.com",
								NickName: "diegof",
							}, true)
							f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "A").Return(true, nil)
							f.teamRepository.Mock.On("DecomposeTeam", mock.Anything, team.Team{
								GroupID: 1,
								UserID:  "A",
							}).Return(nil)
						} else {
							f.userCache.Mock.On("GetByNickName", user).Return(cache.User{
								ID:       "B",
								Name:     "petunia",
								LastName: "avila",
								Email:    "petunia@gmail.com",
								NickName: "petuniaa",
							}, true)
							f.teamRepository.Mock.On("ExistUserInTeam", mock.Anything, "B").Return(true, nil)
							f.teamRepository.Mock.On("DecomposeTeam", mock.Anything, team.Team{
								GroupID: 1,
								UserID:  "B",
							}).Return(nil)
						}
					}
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockTeamService{
				groupRepository: &mocks.IGroupRepository{},
				teamRepository:  &mocks.ITeamRepository{},
				userCache:       &mocks.ICache{},
			}
			tc.mocks.teamService(m)
			service := NewTeamService(m.groupRepository, m.teamRepository, m.userCache)
			err := service.DecomposeTeam(context.Background(), tc.code, tc.teamRequest)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
		})
	}
}
