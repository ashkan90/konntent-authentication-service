package workspaceclient

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"konntent-authentication-service/pkg/httpclient"
	"konntent-authentication-service/pkg/workspaceclient/model"
)

const (
	addWorkspaceEndpoint  = "workspace"
	getWorkspaceEndpoint  = "workspaces/%d"
	getWorkspacesEndpoint = "workspaces"
)

func (c *client) GetWorkspaces(ctx context.Context, req model.GetWorkspacesRequest) ([]model.GetWorkspacesResource, error) {
	res, err := c.HandleRequest(ctx, httpclient.Request{
		URL:     c.PrepareBaseURL(getWorkspacesEndpoint),
		Body:    req,
		Method:  fiber.MethodGet,
		Headers: c.PrepareHeaders(ctx),
	})
	if err != nil {
		return nil, err
	}

	c.logger.Info("GetWorkspaces => ", zap.Any("response", json.RawMessage(res.Body)))

	var resource []model.GetWorkspacesResource
	_ = json.Unmarshal(res.Body, &resource)

	return resource, nil
}

func (c *client) AddWorkspace(ctx context.Context, req model.AddWorkspaceRequest) (*model.AddWorkspaceResource, error) {
	res, err := c.HandleRequest(ctx, httpclient.Request{
		URL:     c.PrepareBaseURL(addWorkspaceEndpoint),
		Method:  fiber.MethodPost,
		Body:    req,
		Headers: c.PrepareHeaders(ctx),
	})
	if err != nil {
		c.logger.Error("an error occurred while making request to external",
			zap.String("error", err.Error()))
		return nil, err
	}

	var resource model.AddWorkspaceResource
	_ = json.Unmarshal(res.Body, &resource)

	return &resource, nil
}
