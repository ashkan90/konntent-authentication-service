package nrclient

import (
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Config struct {
	Key     string
	AppName string
}

type NewRelicInstance interface {
	Application() *newrelic.Application
}

type newRelicInstance struct {
	application *newrelic.Application
}

func InitNewRelic(config Config) (NewRelicInstance, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.AppName),
		newrelic.ConfigLicense(config.Key),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		return nil, err
	}

	return &newRelicInstance{application: app}, nil
}

func (n *newRelicInstance) Application() *newrelic.Application {
	return n.application
}
