//go:build wireinject
// +build wireinject

package konntent_authentication_service

import (
	"go.uber.org/zap"
	"konntent-authentication-service/internal/app"
	"konntent-authentication-service/internal/app/authorize"
	"konntent-authentication-service/internal/app/handler"
	"konntent-authentication-service/internal/app/orchestration"
	"konntent-authentication-service/internal/app/workspace"
	"konntent-authentication-service/pkg/nrclient"
	"konntent-authentication-service/pkg/pg"
	"konntent-authentication-service/pkg/workspaceclient"

	"github.com/google/wire"
)

var serviceProviders = wire.NewSet(
	authorize.NewAuthorizeService,
	workspace.NewWorkspaceService,
)

var repositoryProviders = wire.NewSet(
	authorize.NewAuthorizeRepository,
	workspace.NewWorkspaceRepository,
)

var orchestratorProviders = wire.NewSet(
	orchestration.NewAuthenticationOrchestration,
)

var handlerProviders = wire.NewSet(
	handler.NewAuthHandler,
)

var allProviders = wire.NewSet(
	repositoryProviders,
	serviceProviders,
	orchestratorProviders,
	handlerProviders,
)

func InitAll(
	l *zap.Logger,
	pgInstance pg.Instance,
	workspaceClient workspaceclient.Client,
	nrInstance nrclient.NewRelicInstance,
) app.Router {
	wire.Build(allProviders, app.NewRoute)
	return nil
}
