package team

import (
	"context"

	"cow_back/pkg/platform/database"
	template "cow_back/pkg/platform/templates/team"
	"cow_back/pkg/repository/group"
)

type ITeamRepository interface {
	GetTeamByGroup(ctx context.Context, code string) ([]string, error)
	GetTeamsByUser(ctx context.Context, code string) ([]group.Group, error)
	ExistUserInTeam(ctx context.Context, id string) (bool, error)
	ComposeTeam(ctx context.Context, team Team) (int64, error)
	DecomposeTeam(ctx context.Context, team Team) error
}

type teamRepository struct {
	db database.IDataBase
}

func NewTeamRepository(db database.IDataBase) ITeamRepository {
	return &teamRepository{
		db: db,
	}
}

func (tr *teamRepository) GetTeamByGroup(ctx context.Context, code string) ([]string, error) {
	rs, err := tr.db.GetRead().QueryContext(ctx, template.GetTeamByGroup, code)
	if err != nil {
		return nil, err
	}
	var users []string
	for rs.Next() {
		var user string
		if err := rs.Scan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (tr *teamRepository) GetTeamsByUser(ctx context.Context, code string) ([]group.Group, error) {
	rs, err := tr.db.GetRead().QueryContext(ctx, template.GetTeamsByUser, code)
	if err != nil {
		return nil, err
	}
	var groups []group.Group
	for rs.Next() {
		var group group.Group
		if err := rs.Scan(
			&group.ID,
			&group.Code,
			&group.Debt,
			&group.CreatedAt,
		); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (tr *teamRepository) ExistUserInTeam(ctx context.Context, id string) (bool, error) {
	rs, err := tr.db.GetRead().QueryContext(ctx, template.GetUserByID, id)
	if err != nil {
		return false, err
	}
	rs.Next()
	var exist bool
	if err := rs.Scan(&exist); err != nil {
		return false, err
	}
	return exist, nil
}

func (tr *teamRepository) ComposeTeam(ctx context.Context, team Team) (int64, error) {
	rs, err := tr.db.GetWrite().ExecuteContext(ctx, template.ComposeTeam, team.GroupID, team.UserID)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}

func (tr *teamRepository) DecomposeTeam(ctx context.Context, team Team) error {
	_, err := tr.db.GetWrite().ExecuteContext(ctx, template.DecomposeTeam, team.GroupID, team.UserID)
	if err != nil {
		return err
	}
	return nil
}
