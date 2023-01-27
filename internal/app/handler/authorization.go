package handler

import (
	"github.com/gofiber/fiber/v2"
	"konntent-authentication-service/internal/app/authorize"
	"konntent-authentication-service/internal/app/dto/request"
	"konntent-authentication-service/internal/app/orchestration"
	"konntent-authentication-service/pkg/tokenizer"
	"log"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type authHandler struct {
	authOrchestrator orchestration.Authentication
	repo             authorize.Repository
}

func NewAuthHandler(o orchestration.Authentication, r authorize.Repository) AuthHandler {
	return &authHandler{
		authOrchestrator: o,
		repo:             r,
	}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var req = request.NewLoginInternalRequest(c.Params("uid"))

	// double-check that the uid is valid
	exist := h.repo.CheckID(c.Context(), req.UID)
	if exist != nil {
		log.Println("... ", exist.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	jwtClaim, err := h.authOrchestrator.LoginClaim(c.Context(), req)
	if err != nil {
		log.Println("... ", err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// create new jwt token.
	tokenize := c.Locals("tokenizer").(tokenizer.ITokenizer)
	token := tokenize.Tokenize(tokenizer.NewClaim(jwtClaim)) // Generate Claims struct by somehow

	return c.JSON(token.Token())
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	var (
		req request.Register
		ctx = c.Context()
	)

	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	uid, err := h.authOrchestrator.Register(ctx, req)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	return c.RedirectToRoute("external-login", map[string]interface{}{"uid": uid})
}
