package authorize

import (
	"context"
	"go.uber.org/zap"
	"konntent-authentication-service/internal/app/datamodel"
	"konntent-authentication-service/internal/app/dto/aggregation"
)

type Service interface {
	Login()
	Register(c context.Context, u *datamodel.User) (int, error)
	User(c context.Context, uid int) (*aggregation.User, error)
}

type authorizeService struct {
	repo Repository
	l    *zap.Logger
}

func NewAuthorizeService(l *zap.Logger, r Repository) Service {
	return &authorizeService{
		l:    l,
		repo: r,
	}
}

func (s *authorizeService) Login() {}

func (s *authorizeService) Register(c context.Context, u *datamodel.User) (int, error) {
	return s.repo.AddUser(c, u)
}

func (s *authorizeService) User(c context.Context, uid int) (*aggregation.User, error) {
	var user, err = s.repo.GetUserByID(c, uid)
	if err != nil {
		return nil, err
	}

	return aggregation.NewUserAggregation(user), nil
}
