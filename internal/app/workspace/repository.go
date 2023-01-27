package workspace

import (
	"context"
	pgt "github.com/go-pg/pg/v10"
	"go.uber.org/zap"
	"konntent-authentication-service/internal/app/datamodel"
	"konntent-authentication-service/internal/app/datamodel/resource"
	"konntent-authentication-service/pkg/constants"
	"konntent-authentication-service/pkg/pg"
	pg_conditions "konntent-authentication-service/pkg/pg-conditions"
)

type Repository interface {
	AddUserWorkspace(c context.Context, m datamodel.UserWorkspace) error
	GetUserWorkspaces(c context.Context, m datamodel.Workspace) (*resource.Workspace, error)
}

type workspaceRepository struct {
	pgi pg.Instance
	l   *zap.Logger
}

func NewWorkspaceRepository(l *zap.Logger, pgi pg.Instance) Repository {
	return &workspaceRepository{l: l, pgi: pgi}
}

func (r *workspaceRepository) GetUserWorkspaces(c context.Context, m datamodel.Workspace) (*resource.Workspace, error) {
	var res resource.Workspace
	err := r.pgi.Open().Model(&m).
		Column(constants.GenericIDCol).
		Where(pg_conditions.WhereUserID, m.UserID).
		Select(&res.IDs)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *workspaceRepository) AddUserWorkspace(c context.Context, m datamodel.UserWorkspace) error {
	return r.pgi.Open().RunInTransaction(c, func(tx *pgt.Tx) error {
		_, _wErr := tx.Model(m.ToWorkspace()).Insert()
		if _wErr != nil {
			return _wErr
		}

		_, _err := tx.Model(&m).Insert()
		return _err
	})
}
