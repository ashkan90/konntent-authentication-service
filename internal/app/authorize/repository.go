package authorize

import (
	"context"
	"errors"
	pgt "github.com/go-pg/pg/v10"
	"go.uber.org/zap"
	"konntent-authentication-service/internal/app/datamodel"
	"konntent-authentication-service/pkg/constants"
	"konntent-authentication-service/pkg/pg"
)

type Repository interface {
	AddUser(c context.Context, u *datamodel.User) (int, error)
	GetUserByID(c context.Context, uid int) (*datamodel.User, error)
	CheckID(c context.Context, uid int) error
}

type authorizeRepository struct {
	pgi pg.Instance
	l   *zap.Logger
}

func NewAuthorizeRepository(l *zap.Logger, pgi pg.Instance) Repository {
	return &authorizeRepository{l: l, pgi: pgi}
}

func (r *authorizeRepository) AddUser(c context.Context, u *datamodel.User) (int, error) {
	var err error

	err = r.pgi.Open().RunInTransaction(c, func(tx *pgt.Tx) error {
		_, _err := tx.Model(u).Insert()
		if _err != nil {
			return _err
		}

		u.Account.UserID = u.ID

		_, _err = tx.Model(u.Account).Insert()
		return _err
	})

	if err != nil {
		return constants.UnexpectedUserInsertOccurred, err
	}

	return u.ID, nil
}

func (r *authorizeRepository) GetUserByID(c context.Context, uid int) (*datamodel.User, error) {
	var (
		model = datamodel.User{ID: uid}
		err   error
	)

	err = r.pgi.Open().ModelContext(c, &model).
		Column(constants.UserEmailCol).
		Relation(constants.RelAccountFNameCol).
		Relation(constants.RelAccountLNameCol).
		WherePK().
		First()

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *authorizeRepository) CheckID(c context.Context, uid int) error {
	var i, err = r.pgi.Open().ModelContext(c, &datamodel.User{ID: uid}).WherePK().Count()
	if err != nil {
		return err
	}

	if i == 0 {
		return errors.New("user not found")
	}

	return nil
}
