package sso

import (
	"context"
	"go.uber.org/zap"
	"konntent-authentication-service/pkg/oauth"
	"konntent-authentication-service/pkg/oauthclient/model"
)

type Proxy interface {
	AutoLog(bool, bool)
	Strategy
}

type proxy struct {
	logger   *zap.Logger
	strategy Strategy

	beforeLog bool
	afterLog  bool
}

func NewStrategyProxy(l *zap.Logger, s Strategy) Proxy {
	return &proxy{
		logger:    l,
		strategy:  s,
		beforeLog: true,
	}
}

func (p *proxy) AutoLog(before, after bool) {
	p.beforeLog = before
	p.afterLog = after
}

func (p *proxy) Login() (string, error) {
	p.logB("before")
	defer p.logA("after")

	return p.strategy.Login()
}

func (p *proxy) Register(c context.Context, r oauth.CallbackResponse) (*model.Generic, error) {
	p.logB("before")
	defer p.logA("after")
	return p.strategy.Register(c, r)
}

func (p *proxy) String() string {
	return p.strategy.String()
}

func (p *proxy) logB(content string) {
	if p.beforeLog {
		p.logger.Info(content)
	}
}

func (p *proxy) logA(content string) {
	if p.afterLog {
		p.logger.Info(content)
	}
}
