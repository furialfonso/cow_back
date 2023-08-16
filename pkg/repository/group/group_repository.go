package group

import (
	"context"
	"docker-go-project/pkg/platform/database"

	template "docker-go-project/pkg/platform/templates/group"
	"fmt"
)

type IGroupRepository interface {
	GetAll(ctx context.Context) ([]Group, error)
	GetByCode(ctx context.Context, code string) (Group, error)
	Create(ctx context.Context, code string) (int64, error)
	Delete(ctx context.Context, code string) error
	UpdateDebtByCode(ctx context.Context, group Group) error
}

type groupRepository struct {
	db database.IDataBase
}

func NewGroupRepository(db database.IDataBase) IGroupRepository {
	return &groupRepository{
		db: db,
	}
}

func (gr *groupRepository) GetAll(ctx context.Context) ([]Group, error) {
	var groups []Group
	rs, err := gr.db.GetRead().QueryContext(ctx, template.GetAll)
	if err != nil {
		return groups, err
	}
	for rs.Next() {
		var group Group
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

func (gr *groupRepository) GetByCode(ctx context.Context, code string) (Group, error) {
	var group Group
	rs, err := gr.db.GetRead().QueryContext(ctx, template.GetByCode, code)
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

func (gr *groupRepository) Create(ctx context.Context, code string) (int64, error) {
	rs, err := gr.db.GetWrite().ExecuteContext(ctx, template.Create, code)
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
	_, err := gr.db.GetWrite().ExecuteContext(ctx, template.Delete, code)
	if err != nil {
		return err
	}
	return nil
}

func (gr *groupRepository) UpdateDebtByCode(ctx context.Context, group Group) error {
	_, err := gr.db.GetWrite().ExecuteContext(ctx, template.Update, group.Debt, group.Code)
	if err != nil {
		return err
	}
	return nil
}
