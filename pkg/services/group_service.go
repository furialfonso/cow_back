package services

import (
	"context"
	"docker-go-project/api/dto/request"
	"docker-go-project/api/dto/response"
	"docker-go-project/pkg/repository"
	"docker-go-project/pkg/repository/model"
)

type IGroupService interface {
	GetGroups(ctx context.Context) ([]response.GroupResponse, error)
	GetGroupByCode(ctx context.Context, code string) (response.GroupResponse, error)
	CreateGroup(ctx context.Context, groupDTO request.GroupDTO) error
	UpdateDebtByCode(ctx context.Context, groupDTO request.GroupDTO) error
}

type groupService struct {
	groupRepository repository.IGroupRepository
}

func NewGroupService(groupRepository repository.IGroupRepository) IGroupService {
	return &groupService{
		groupRepository: groupRepository,
	}
}

func (gs *groupService) GetGroups(ctx context.Context) ([]response.GroupResponse, error) {
	var groupsResponse []response.GroupResponse
	groups, err := gs.groupRepository.GetGroups(ctx)
	if err != nil {
		return groupsResponse, err
	}
	for _, group := range groups {
		groupsResponse = append(groupsResponse, response.GroupResponse{
			Code: group.Code,
		})
	}
	return groupsResponse, nil
}

func (gs *groupService) GetGroupByCode(ctx context.Context, code string) (response.GroupResponse, error) {
	var groupResponse response.GroupResponse
	rs, err := gs.groupRepository.GetGroupByCode(ctx, code)
	if err != nil {
		return groupResponse, err
	}
	groupResponse.Code = rs.Code
	return groupResponse, nil
}

func (gs *groupService) CreateGroup(ctx context.Context, groupDTO request.GroupDTO) error {
	_, err := gs.groupRepository.Create(ctx, groupDTO.Code)
	if err != nil {
		return err
	}
	return nil
}

func (gs *groupService) UpdateDebtByCode(ctx context.Context, groupDTO request.GroupDTO) error {
	err := gs.groupRepository.UpdateGroupDebtByCode(ctx, model.Group{
		Code: groupDTO.Code,
		Debt: groupDTO.Debt,
	})

	if err != nil {
		return err
	}
	return nil
}
