package group

import (
	"context"
	"errors"
	"testing"

	"cow_back/api/dto/request"
	"cow_back/api/dto/response"
	"cow_back/mocks"
	"cow_back/pkg/repository/group"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGroupService struct {
	groupRepository *mocks.IGroupRepository
}

type groupMocks struct {
	groupService func(f *mockGroupService)
}

func Test_GetAll(t *testing.T) {
	tests := []struct {
		expErr error
		mocks  groupMocks
		name   string
		outPut []response.GroupResponse
	}{
		{
			name: "error get groups",
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("GetAll", mock.Anything).Return([]group.Group{}, errors.New("error x"))
				},
			},
			expErr: errors.New("error x"),
		},
		{
			name: "full flow",
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("GetAll", mock.Anything).Return([]group.Group{
						{
							ID:        1,
							Code:      "YOU&I",
							Debt:      200000,
							CreatedAt: "2023-05-01T08:00:00",
						},
						{
							ID:        2,
							Code:      "test1",
							Debt:      300000,
							CreatedAt: "2023-05-01T08:00:00",
						},
					}, nil)
				},
			},
			outPut: []response.GroupResponse{
				{
					Code: "YOU&I",
				},
				{
					Code: "test1",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockGroupService{
				groupRepository: &mocks.IGroupRepository{},
			}
			tc.mocks.groupService(m)
			service := NewGroupService(m.groupRepository)
			groups, err := service.GetAll(context.Background())
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.outPut, groups)
		})
	}
}

func Test_GetByCode(t *testing.T) {
	tests := []struct {
		expErr error
		mocks  groupMocks
		name   string
		code   string
		outPut response.GroupResponse
	}{
		{
			name: "error get groups",
			code: "YOU&I",
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "YOU&I").Return(group.Group{}, errors.New("error x"))
				},
			},
			expErr: errors.New("error x"),
		},
		{
			name: "full flow",
			code: "YOU&I",
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("GetByCode", mock.Anything, "YOU&I").Return(group.Group{
						ID:        1,
						Code:      "YOU&I",
						Debt:      200000,
						CreatedAt: "2023-05-01T08:00:00",
					}, nil)
				},
			},
			outPut: response.GroupResponse{
				Code: "YOU&I",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockGroupService{
				groupRepository: &mocks.IGroupRepository{},
			}
			tc.mocks.groupService(m)
			service := NewGroupService(m.groupRepository)
			groups, err := service.GetByCode(context.Background(), tc.code)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.outPut, groups)
		})
	}
}

func Test_Create(t *testing.T) {
	tests := []struct {
		expErr error
		mocks  groupMocks
		name   string
		input  request.GroupRequest
	}{
		{
			name: "error",
			input: request.GroupRequest{
				Code: "Test1",
			},
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("Create", mock.Anything, "Test1").Return(int64(0), errors.New("error x"))
				},
			},
			expErr: errors.New("error x"),
		},
		{
			name: "full flow",
			input: request.GroupRequest{
				Code: "Test1",
			},
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("Create", mock.Anything, "Test1").Return(int64(1), nil)
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockGroupService{
				groupRepository: &mocks.IGroupRepository{},
			}
			tc.mocks.groupService(m)
			service := NewGroupService(m.groupRepository)
			err := service.Create(context.Background(), tc.input)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
		})
	}
}

func Test_Delete(t *testing.T) {
	tests := []struct {
		expErr error
		mocks  groupMocks
		name   string
		input  string
	}{
		{
			name:  "error",
			input: "Test1",
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("Delete", mock.Anything, "Test1").Return(errors.New("error x"))
				},
			},
			expErr: errors.New("error x"),
		},
		{
			name:  "full flow",
			input: "Test1",
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("Delete", mock.Anything, "Test1").Return(nil)
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockGroupService{
				groupRepository: &mocks.IGroupRepository{},
			}
			tc.mocks.groupService(m)
			service := NewGroupService(m.groupRepository)
			err := service.Delete(context.Background(), tc.input)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
		})
	}
}

func Test_UpdateDebtByCode(t *testing.T) {
	tests := []struct {
		expErr error
		mocks  groupMocks
		name   string
		input  request.GroupRequest
	}{
		{
			name: "error",
			input: request.GroupRequest{
				Code: "Test1",
				Debt: 1000,
			},
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("UpdateDebtByCode", mock.Anything, group.Group{
						Code: "Test1",
						Debt: 1000,
					}).Return(errors.New("error x"))
				},
			},
			expErr: errors.New("error x"),
		},
		{
			name: "full flow",
			input: request.GroupRequest{
				Code: "Test1",
				Debt: 1000,
			},
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("UpdateDebtByCode", mock.Anything, group.Group{
						Code: "Test1",
						Debt: 1000,
					}).Return(nil)
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockGroupService{
				groupRepository: &mocks.IGroupRepository{},
			}
			tc.mocks.groupService(m)
			service := NewGroupService(m.groupRepository)
			err := service.UpdateDebtByCode(context.Background(), tc.input)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
		})
	}
}
