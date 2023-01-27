package sso_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"konntent-authentication-service/pkg/sso"
	"konntent-authentication-service/pkg/sso/strategies"
)

var _ = Describe("SSO Unit Tests", func() {
	Describe("init sso with given correct strategy/proxy", func() {
		It("should return new sso selector", func() {
			_sso := sso.InitSSO(sso.NewStrategyProxy(
				zap.L(),
				strategies.NewGithubSSO(),
			))
			Expect(_sso).To(Not(BeNil()))
		})
	})
})
