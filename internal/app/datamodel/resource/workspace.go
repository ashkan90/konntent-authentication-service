package resource

import "konntent-authentication-service/pkg/workspaceclient/model"

type Workspace struct {
	IDs []int
}

func (w Workspace) ToModel() model.GetWorkspacesRequest {
	return model.GetWorkspacesRequest{
		WorkspaceIDs: w.IDs,
	}
}
