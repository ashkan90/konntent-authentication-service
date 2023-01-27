package model

type AddWorkspaceRequest struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	BM          string `json:"businessModel"`
	CompanyUnit int    `json:"companyUnit"`
	UserID      int    `json:"userId"`
}

type AddWorkspaceResource struct {
	WorkspaceID int `json:"workspaceId"`
}

type GetWorkspacesByUIDRequest struct {
	UserID int `json:"userId"`
}

type GetWorkspacesRequest struct {
	WorkspaceIDs []int `json:"workspaceIds"`
}

type GetWorkspacesResource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
