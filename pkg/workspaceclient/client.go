package workspaceclient

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"konntent-authentication-service/pkg/httpclient"
	"konntent-authentication-service/pkg/workspaceclient/model"
)

type Client interface {
	HandleRequest(ctx context.Context, req httpclient.Request) (*httpclient.Response, error)
	PrepareHeaders(ctx context.Context) map[string]string
	PrepareBaseURL(u string) string

	AddWorkspace(ctx context.Context, req model.AddWorkspaceRequest) (*model.AddWorkspaceResource, error)
	GetWorkspaces(ctx context.Context, req model.GetWorkspacesRequest) ([]model.GetWorkspacesResource, error)
}

type Config struct {
	BaseURL string
	Timeout int
}

type client struct {
	baseURL    string
	timeout    int
	httpClient httpclient.HTTPClient
	logger     *zap.Logger
}

func NewClient(l *zap.Logger, c Config, hc httpclient.HTTPClient) Client {
	return &client{
		baseURL:    c.BaseURL,
		timeout:    c.Timeout,
		httpClient: hc,
		logger:     l,
	}
}

func (c *client) HandleRequest(ctx context.Context, req httpclient.Request) (*httpclient.Response, error) {
	resp, err := c.httpClient.HandleRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	if c.httpClient.IsSuccessStatusCode(resp) {
		return resp, nil
	}

	return nil, c.httpClient.HandleException(resp)
}

func (c *client) PrepareHeaders(ctx context.Context) map[string]string {
	headers := c.httpClient.GetJSONHeaders()

	return headers
}

func (c *client) PrepareBaseURL(u string) string {
	return fmt.Sprintf("%s/%s", c.baseURL, u)
}
