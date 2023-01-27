package oauthclient

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"konntent-authentication-service/pkg/httpclient"
	"konntent-authentication-service/pkg/oauthclient/model"
)

func (c *client) UserInfoGithub(ctx context.Context, token string) (*model.Generic, error) {
	var resp, err = c.HandleRequest(ctx, httpclient.Request{
		URL: "https://api.github.com/user",
		Headers: map[string]string{
			"Authorization": "token " + token,
		},
		Method: fiber.MethodGet,
	})
	if err != nil {
		return nil, err
	}

	var resource model.Generic
	_ = json.Unmarshal(resp.Body, &resource)

	return &resource, err
}
