package aggregation

import (
	"konntent-authentication-service/pkg/workspaceclient/model"
)

type Workspace struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func NewWorkspacesAggregation(w []model.GetWorkspacesResource) []Workspace {
	var ws = make([]Workspace, 0, len(w))

	for _, resource := range w {
		ws = append(ws, Workspace{
			ID:   resource.ID,
			Name: resource.Name,
			URL:  resource.URL,
		})
	}

	return ws
}
