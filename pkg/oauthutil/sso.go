package oauthutil

import (
	"context"
	"konntent-authentication-service/pkg/oauth"
	"konntent-authentication-service/pkg/sso"
	"konntent-authentication-service/pkg/utils"
)

func Google(c context.Context) sso.Proxy {
	return ByProvider(c, oauth.ProviderGoogle)
}

func Github(c context.Context) sso.Proxy {
	return ByProvider(c, oauth.ProviderGithub)
}

func ByProvider(c context.Context, provider string) sso.Proxy {
	return c.Value(utils.Providers[provider]).(sso.Proxy)
}
