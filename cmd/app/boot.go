package main

import (
	"fmt"
	"go.uber.org/zap"
	"konntent-authentication-service/configs/app"
	"konntent-authentication-service/pkg/constants"
	"konntent-authentication-service/pkg/httpclient"
	"konntent-authentication-service/pkg/nrclient"
	"konntent-authentication-service/pkg/pg"
	pg_migration "konntent-authentication-service/pkg/pg-migration"
	pg_rel_registration "konntent-authentication-service/pkg/pg-rel-registration"
	"konntent-authentication-service/pkg/workspaceclient"
	"time"

	"github.com/spf13/viper"
)

func boot(l *zap.Logger, appConf app.ApplicationConfigs) (*server, error) {
	time.Local, _ = time.LoadLocation("Europe/Istanbul")

	var pgInstance, err = initPG(l, appConf.Postgres)
	if err != nil {
		return nil, err
	}

	var (
		httpClient      = httpclient.NewHTTPClient()
		workspaceClient = workspaceclient.NewClient(l, initWorkspaceClientConfig(appConf.Clients.Workspace), httpClient)
	)

	//var (
	//oauthClient = oauthclient.NewClient(httpClient)
	//oauthGoogleProcessor = oauth.NewOAuthProcessor(oauthClient, appConf.Auth.Google)
	//oauthGoogle          = sso.NewStrategyProxy(zap.L(), strategies.NewGoogleSSO(oauthGoogleProcessor))
	//oauthSSO             = sso.InitSSO(oauthGoogle)
	//)
	//err := mqProducer.ConnectToBroker(appConf.Rabbit.URL)
	//if err != nil {
	//	return nil, err
	//}

	//nrInstance, err := initNewRelic(appConf.NewRelic)
	//if err != nil {
	//	return nil, err
	//}

	return &server{
		logger:          l,
		pgInstance:      pgInstance,
		workspaceClient: workspaceClient,
		tokenizerConfig: appConf.JWTConfig,
		//oauth:       oauthSSO,
		//oauthClient: oauthClient,
		//oauthConfig: appConf.Auth,
		//nrInstance:  nrInstance,
	}, nil
}

func initWorkspaceClientConfig(cfg app.ClientConfig) workspaceclient.Config {
	return workspaceclient.Config{
		BaseURL: cfg.BaseURL,
		Timeout: cfg.Timeout,
	}
}

func initNewRelic(cfg app.NewRelicConfig) (nrclient.NewRelicInstance, error) {
	return nrclient.InitNewRelic(nrclient.Config{
		Key:     cfg.ApplicationKey,
		AppName: cfg.ApplicationName,
	})
}

func initPG(l *zap.Logger, cfg app.PGSettings) (pg.Instance, error) {
	return pg.NewPGInstance(l, cfg)
}

func initConfig[T string | constants.AppEnvironment](env T) (*app.Configs, error) {
	if env == "" {
		env = T(constants.ConfigEnvDefault)
	}
	viper.AutomaticEnv()
	viper.SetConfigName(fmt.Sprintf("%s.%s", env, "server"))
	viper.SetConfigType(constants.ConfigEnvFileType)
	viper.AddConfigPath(constants.ConfigEnvFilePath)
	viper.AddConfigPath(constants.ConfigEnvFilePathContainer)

	var appConf app.Configs

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&appConf)
	if err != nil {
		return nil, err
	}

	//if strings.HasPrefix(appConf.Application.Rabbit.URL, "$") {
	//	appConf.Application.Rabbit.URL = viper.GetString(constants.ConfigAMQPEnvKey)
	//}

	appConf.Application.NewRelic = app.NewRelicConfig{
		ApplicationKey:  viper.GetString(constants.ConfigNRLicenseKey),
		ApplicationName: viper.GetString(constants.ConfigNRAppKey),
	}

	return &appConf, nil
}

func registrar(l *zap.Logger) {
	pg_rel_registration.Register()
}

func migrate(l *zap.Logger, pg pg.Instance) {
	err := pg_migration.Migrate(pg, pg_migration.MigrationModels...)
	if err != nil {
		l.Error("something went wrong while migrating...",
			zap.Error(err))
	}
}
