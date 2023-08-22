package team

import (
	"context"
	"docker-go-project/pkg/platform/database"
	template "docker-go-project/pkg/platform/templates/team"
)

type ITeamRepository interface {
	GetUsersByGroup(ctx context.Context, code string) ([]string, error)
	ExistUserInTeam(ctx context.Context, id int64) (bool, error)
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

func (tr *teamRepository) GetUsersByGroup(ctx context.Context, code string) ([]string, error) {
	rs, err := tr.db.GetRead().QueryContext(ctx, template.GetUsersByGroup, code)
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

func (tr *teamRepository) ExistUserInTeam(ctx context.Context, id int64) (bool, error) {
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
