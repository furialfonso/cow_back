package repository

import (
	"context"
	"docker-go-project/pkg/platform/database"
	"docker-go-project/pkg/platform/sql/templates"
	"docker-go-project/pkg/repository/model"
	"fmt"
)

type IGroupRepository interface {
	GetGroupByCode(ctx context.Context, code string) (model.Group, error)
	GetGroups(ctx context.Context) ([]model.Group, error)
	Create(ctx context.Context, code string) (int64, error)
	Delete(ctx context.Context, code string) error
	UpdateGroupDebtByCode(ctx context.Context, group model.Group) error
}

type groupRepository struct {
	db database.IDataBase
}

func NewRepository(db database.IDataBase) IGroupRepository {
	return &groupRepository{
		db: db,
	}
}

func (gr *groupRepository) GetGroupByCode(ctx context.Context, code string) (model.Group, error) {
	var group model.Group
	rs, err := gr.db.GetRead().QueryContext(ctx, templates.GetGroupByCode, code)
	if err != nil {
		return group, err
	}
	exists := rs.Next()
	if !exists {
		return group, fmt.Errorf("group %s not found", code)
	}
	if err := rs.Scan(
		&group.ID,
		&group.Code,
		&group.Debt,
		&group.CreatedAt,
	); err != nil {
		return group, err
	}
	return group, nil
}

func (gr *groupRepository) GetGroups(ctx context.Context) ([]model.Group, error) {
	var groups []model.Group
	rs, err := gr.db.GetRead().QueryContext(ctx, templates.GetGroups)
	if err != nil {
		return groups, err
	}
	for rs.Next() {
		var group model.Group
		if err := rs.Scan(
			&group.ID,
			&group.Code,
			&group.Debt,
			&group.CreatedAt,
		); err != nil {
			return groups, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (gr *groupRepository) Create(ctx context.Context, code string) (int64, error) {
	rs, err := gr.db.GetWrite().ExecuteContext(ctx, templates.CreateGroup, code)
	if err != nil {
		return 0, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (gr *groupRepository) Delete(ctx context.Context, code string) error {
	_, err := gr.db.GetWrite().ExecuteContext(ctx, templates.DeleteGroup, code)
	if err != nil {
		return err
	}
	return nil
}

func (gr *groupRepository) UpdateGroupDebtByCode(ctx context.Context, group model.Group) error {
	_, err := gr.db.GetWrite().ExecuteContext(ctx, templates.UpdateGroup, group.Debt, group.Code)
	if err != nil {
		return err
	}
	return nil
}
