package strategies

import (
	"context"
	"konntent-authentication-service/pkg/oauth"
	"konntent-authentication-service/pkg/oauthclient/model"
	"konntent-authentication-service/pkg/sso"
)

type Github struct {
	processor *oauth.Processor
}

func NewGithubSSO(p *oauth.Processor) sso.Strategy {
	return &Github{processor: p}
}

func (g *Github) Login() (string, error) {
	return g.processor.BuildLoginURL()
}

func (g *Github) Register(c context.Context, r oauth.CallbackResponse) (*model.Generic, error) {
	var tok, err = g.processor.GetExchange(c, r)
	if err != nil {
		return nil, err
	}

	var res, reqErr = g.processor.GetUserInfo(c, tok)
	if reqErr != nil {
		return nil, reqErr
	}

	return res, reqErr
}

func (g *Github) String() string {
	return "github-sso"
}
