package app

import (
	"github.com/gofiber/fiber/v2"
	"konntent-authentication-service/internal/app/handler"
)

type RouteCtx struct {
	App *fiber.App
}

type Router interface {
	SetupRoutes(r *RouteCtx)
}

type route struct {
	authHandler handler.AuthHandler
}

func NewRoute(authHandler handler.AuthHandler) Router {
	return &route{
		authHandler: authHandler,
	}
}

func (r *route) SetupRoutes(rc *RouteCtx) {
	rc.App.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	rc.App.Post("/register", r.authHandler.Register)
	loginGroup := rc.App.Group("/login")

	r.loginRoutes(loginGroup)

	rc.App.Post("/login", r.authHandler.Login)
	rc.App.Post("/login/:uid/external", r.authHandler.Login).Name("external-login")
}

func (r *route) loginRoutes(gr fiber.Router) {
	gr.Get("/", r.authHandler.Login)
}
