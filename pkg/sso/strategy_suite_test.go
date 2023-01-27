package sso_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSSO(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SSO Unit Tests")
}
