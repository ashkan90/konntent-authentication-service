package orchestration

import (
	"context"
	"go.uber.org/zap"
	"konntent-authentication-service/internal/app/authorize"
	"konntent-authentication-service/internal/app/dto/aggregation"
	"konntent-authentication-service/internal/app/dto/request"
	"konntent-authentication-service/internal/app/workspace"
)

type Authentication interface {
	LoginClaim(c context.Context, dto request.LoginInternalRequest) (authorize.JWTClaim, error)
	Register(c context.Context, u request.Register) (int, error)
}

type authenticationOrchestration struct {
	l                *zap.Logger
	authService      authorize.Service
	workspaceService workspace.Service
}

func NewAuthenticationOrchestration(l *zap.Logger, s authorize.Service, ws workspace.Service) Authentication {
	return &authenticationOrchestration{
		l:                l,
		authService:      s,
		workspaceService: ws,
	}
}

func (ao *authenticationOrchestration) LoginClaim(c context.Context, dto request.LoginInternalRequest) (authorize.JWTClaim, error) {
	var userAggr, uErr = ao.authService.User(c, dto.UID)
	if uErr != nil {
		return authorize.JWTClaim{}, uErr
	}

	var workspaces, wErr = ao.workspaceService.GetWorkspaces(c, dto.ToModel())
	if wErr != nil {
		return authorize.JWTClaim{}, wErr
	}

	aggr := authorize.NewClaimAggregation(userAggr, aggregation.NewWorkspacesAggregation(workspaces))

	return aggr, nil
}

func (ao *authenticationOrchestration) Register(c context.Context, dto request.Register) (int, error) {
	var (
		uid int
		err error
	)

	uid, err = ao.authService.Register(c, dto.User.ToUserModel())
	if err != nil {
		return uid, err
	}

	err = ao.workspaceService.AddWorkspace(c, dto.Workspace.ToModel(uid))
	if err != nil {
		return uid, err
	}

	return uid, nil
}
