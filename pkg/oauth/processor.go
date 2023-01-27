package oauth

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"konntent-authentication-service/configs/app"
	"konntent-authentication-service/pkg/oauthclient"
	"konntent-authentication-service/pkg/oauthclient/model"
	neturl "net/url"
	"strings"
)

type Processor struct {
	client oauthclient.Client
	conf   *oauth2.Config
}

func NewOAuthProcessor(client oauthclient.Client, conf app.GeneralOAuthSettings) *Processor {
	var endpoint = google.Endpoint
	if strings.HasSuffix(conf.RedirectURL, "github") {
		endpoint = github.Endpoint
	}

	return &Processor{
		client: client,
		conf: &oauth2.Config{
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			RedirectURL:  conf.RedirectURL,
			Scopes:       conf.Scopes,
			Endpoint:     endpoint,
		},
	}
}

func (p Processor) BuildLoginURL() (string, error) {
	var (
		url, err = neturl.Parse(p.conf.Endpoint.AuthURL)
		params   = neturl.Values{
			LoginCredentialClientID:     []string{p.conf.ClientID},
			LoginCredentialRedirectURI:  []string{p.conf.RedirectURL},
			LoginCredentialResponseType: []string{"code"},
			LoginCredentialState:        []string{"stateString"},
			LoginCredentialScope:        []string{strings.Join(p.conf.Scopes, " ")},
		}
	)

	if err != nil {
		return "", err
	}

	url.RawQuery = params.Encode()

	return url.String(), nil
}

func (p Processor) GetExchange(c context.Context, resp CallbackResponse) (*oauth2.Token, error) {
	var tok, err = p.conf.Exchange(c, resp.Code)
	if err != nil {
		return nil, err
	}

	return tok, nil
}

func (p Processor) GetUserInfo(c context.Context, tok *oauth2.Token) (*model.Generic, error) {
	var (
		res *model.Generic
		err error
	)

	if strings.HasSuffix(p.conf.RedirectURL, "github") {
		res, err = p.client.UserInfoGithub(c, tok.AccessToken)
	} else {
		res, err = p.client.UserInfoGoogle(c, tok.AccessToken)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
