package group

import (
	"context"

	"cow_back/api/dto/request"
	"cow_back/api/dto/response"
	"cow_back/pkg/repository/group"
)

type IGroupService interface {
	GetAll(ctx context.Context) ([]response.GroupResponse, error)
	GetByCode(ctx context.Context, code string) (response.GroupResponse, error)
	Create(ctx context.Context, groupRequest request.GroupRequest) error
	Delete(ctx context.Context, code string) error
	UpdateDebtByCode(ctx context.Context, groupRequest request.GroupRequest) error
}

type groupService struct {
	groupRepository group.IGroupRepository
}

func NewGroupService(groupRepository group.IGroupRepository) IGroupService {
	return &groupService{
		groupRepository: groupRepository,
	}
}

func (gs *groupService) GetAll(ctx context.Context) ([]response.GroupResponse, error) {
	var groupsResponse []response.GroupResponse
	groups, err := gs.groupRepository.GetAll(ctx)
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

func (gs *groupService) GetByCode(ctx context.Context, code string) (response.GroupResponse, error) {
	var groupResponse response.GroupResponse
	rs, err := gs.groupRepository.GetByCode(ctx, code)
	if err != nil {
		return groupResponse, err
	}
	groupResponse.Code = rs.Code
	return groupResponse, nil
}

func (gs *groupService) Create(ctx context.Context, groupRequest request.GroupRequest) error {
	_, err := gs.groupRepository.Create(ctx, groupRequest.Code)
	if err != nil {
		return err
	}
	return nil
}

func (gs *groupService) Delete(ctx context.Context, code string) error {
	err := gs.groupRepository.Delete(ctx, code)
	if err != nil {
		return err
	}
	return nil
}

func (gs *groupService) UpdateDebtByCode(ctx context.Context, groupRequest request.GroupRequest) error {
	err := gs.groupRepository.UpdateDebtByCode(ctx, group.Group{
		Code: groupRequest.Code,
		Debt: groupRequest.Debt,
	})
	if err != nil {
		return err
	}
	return nil
}
