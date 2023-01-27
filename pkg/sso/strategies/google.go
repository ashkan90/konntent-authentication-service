package strategies

import (
	"context"
	"konntent-authentication-service/pkg/oauth"
	"konntent-authentication-service/pkg/oauthclient/model"
	"konntent-authentication-service/pkg/sso"
)

type Google struct {
	processor *oauth.Processor
}

func NewGoogleSSO(p *oauth.Processor) sso.Strategy {
	return &Google{processor: p}
}

func (g *Google) Login() (string, error) {
	return g.processor.BuildLoginURL()
}

func (g *Google) Register(c context.Context, r oauth.CallbackResponse) (*model.Generic, error) {
	var tok, err = g.processor.GetExchange(c, r)
	if err != nil {
		return nil, err
	}

	var res, reqErr = g.processor.GetUserInfo(c, tok)

	return res, reqErr
}

func (g *Google) String() string {
	return "google-sso"
}
