package workspace

import (
	"context"
	"go.uber.org/zap"
	dm "konntent-authentication-service/internal/app/datamodel"
	"konntent-authentication-service/pkg/workspaceclient"
	"konntent-authentication-service/pkg/workspaceclient/model"
)

type Service interface {
	AddWorkspace(c context.Context, req model.AddWorkspaceRequest) error
	GetWorkspaces(c context.Context, req model.GetWorkspacesByUIDRequest) ([]model.GetWorkspacesResource, error)
}

type workspaceService struct {
	l                   *zap.Logger
	workspaceClient     workspaceclient.Client
	workspaceRepository Repository
}

func NewWorkspaceService(l *zap.Logger, wr Repository, wc workspaceclient.Client) Service {
	return &workspaceService{
		l:                   l,
		workspaceClient:     wc,
		workspaceRepository: wr,
	}
}

func (w *workspaceService) GetWorkspaces(c context.Context, req model.GetWorkspacesByUIDRequest) ([]model.GetWorkspacesResource, error) {
	var res, _err = w.workspaceRepository.GetUserWorkspaces(c, dm.NewWorkspaceModel(req.UserID))
	if _err != nil {
		return nil, _err
	}

	var workspaces, err = w.workspaceClient.GetWorkspaces(c, res.ToModel())
	if err != nil {
		return nil, err
	}

	return workspaces, err
}

func (w *workspaceService) AddWorkspace(c context.Context, req model.AddWorkspaceRequest) error {
	var res, err = w.workspaceClient.AddWorkspace(c, req)
	if err != nil {
		return err
	}

	err = w.workspaceRepository.AddUserWorkspace(c, dm.NewUserWorkspaceModel(req.UserID, res.WorkspaceID))
	if err != nil {
		return err
	}

	return nil
}
