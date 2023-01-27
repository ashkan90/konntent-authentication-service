package sso

import (
	"context"
	"konntent-authentication-service/pkg/utils"
)

func ProxyByContext(c context.Context, p string) Proxy {
	return c.Value(utils.Providers[p]).(Proxy)
}
