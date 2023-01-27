package middlewarepkg

import (
	"konntent-authentication-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func PutHeaders(ctx *fiber.Ctx) error {
	var headerMap = make(map[string]string)

	ctx.Request().Header.VisitAll(func(key, value []byte) {
		if string(key) == utils.HeaderAuthorization {
			headerMap[string(key)] = string(value)
		}
	})

	ctx.Locals(utils.HeaderMapCtx, headerMap)
	return ctx.Next()
}

func GetAuthorizationHeader(ctx *fiber.Ctx) []byte {
	var headerMap = ctx.Locals(utils.HeaderMapCtx).(map[string]string)

	if val, ok := headerMap[utils.HeaderAuthorization]; ok {
		return []byte(val)
	}

	return []byte{}
}
