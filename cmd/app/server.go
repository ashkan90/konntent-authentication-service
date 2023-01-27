package main

import (
	"github.com/gofiber/fiber/v2"
	recoverpkg "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	di "konntent-authentication-service"
	configs "konntent-authentication-service/configs/app"
	"konntent-authentication-service/internal/app"
	"konntent-authentication-service/internal/app/middleware"
	"konntent-authentication-service/pkg/constants"
	"konntent-authentication-service/pkg/middlewarepkg"
	"konntent-authentication-service/pkg/nrclient"
	"konntent-authentication-service/pkg/oauthclient"
	"konntent-authentication-service/pkg/pg"
	"konntent-authentication-service/pkg/sso"
	"konntent-authentication-service/pkg/tokenizer"
	"konntent-authentication-service/pkg/utils"
	"konntent-authentication-service/pkg/validation"
	"konntent-authentication-service/pkg/workspaceclient"
)

type server struct {
	logger          *zap.Logger
	pgInstance      pg.Instance
	workspaceClient workspaceclient.Client
	tokenizerConfig configs.JWTConfig

	oauth       sso.StrategySelector
	oauthClient oauthclient.Client
	oauthConfig configs.AuthConfig
	nrInstance  nrclient.NewRelicInstance
}

func initServer(sv *server) *fiber.App {
	fApp := fiber.New(fiber.Config{
		BodyLimit: constants.AppRequestBodyLimit,
	})
	fApp.Use(recoverpkg.New(recoverpkg.Config{
		EnableStackTrace: true,
	}))

	sv.initCommonMiddlewares(fApp)

	route := di.InitAll(
		sv.logger,
		sv.pgInstance,
		sv.workspaceClient,
		sv.nrInstance,
	)
	route.SetupRoutes(&app.RouteCtx{
		App: fApp,
	})

	return fApp
}

func initLogger() *zap.Logger {
	zc := zap.NewDevelopmentEncoderConfig()
	zc.EncodeLevel = zapcore.CapitalColorLevelEncoder

	l := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zc),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	))
	return l
}

func (s *server) initCommonMiddlewares(app *fiber.App) {
	validator := validation.InitValidator()
	tokenize := tokenizer.NewTokenizer(s.tokenizerConfig)

	app.Use(middleware.NewRelicMiddleware(s.nrInstance))
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(utils.Validator, validator)
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("tokenizer", tokenize)
		return c.Next()
	})

	//app.Use(func(c *fiber.Ctx) error {
	//	c.Locals(utils.Claimer, s.jwtInstance)
	//	return c.Next()
	//})

	//app.Use(func(c *fiber.Ctx) error {
	//	var (
	//		oauthGoogleProcessor = oauth.NewOAuthProcessor(s.oauthClient, s.oauthConfig.Google)
	//		oauthGoogle          = sso.NewStrategyProxy(s.logger, strategies.NewGoogleSSO(oauthGoogleProcessor))
	//
	//		oauthGithubProcessor = oauth.NewOAuthProcessor(s.oauthClient, s.oauthConfig.Github)
	//		oauthGithub          = sso.NewStrategyProxy(s.logger, strategies.NewGithubSSO(oauthGithubProcessor))
	//	)
	//
	//	c.Locals(oauth.GithubCtx, oauthGithub)
	//	c.Locals(oauth.GoogleCtx, oauthGoogle)
	//
	//	return c.Next()
	//})

	app.Use(middlewarepkg.PutHeaders)
}
