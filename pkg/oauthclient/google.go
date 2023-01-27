package oauthclient

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"konntent-authentication-service/pkg/httpclient"
	"konntent-authentication-service/pkg/oauthclient/model"
)

func (c *client) UserInfoGoogle(ctx context.Context, token string) (*model.Generic, error) {
	var resp, err = c.HandleRequest(ctx, httpclient.Request{
		URL:    prepareURL(token),
		Method: fiber.MethodGet,
	})

	if err != nil {
		return nil, err
	}

	var resource model.Generic
	_ = json.Unmarshal(resp.Body, &resource)

	resource.Token = token

	return &resource, nil
}

func prepareURL(token string) string {
	return "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token
}
