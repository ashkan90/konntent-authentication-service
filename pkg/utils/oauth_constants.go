package utils

import "konntent-authentication-service/pkg/oauth"

var (
	Providers = map[string]string{
		oauth.ProviderGithub: oauth.GithubCtx,
		oauth.ProviderGoogle: oauth.GoogleCtx,
	}
)
