package claimer

import (
	"context"
	"encoding/json"
	"konntent-authentication-service/pkg/utils"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Claimer interface {
	IsValid(ctx context.Context, rawToken []byte) ([]byte, bool)
	GetModel(ctx context.Context) *Model
}

type Model struct {
}

type claimer struct {
	signKey string
}

func NewClaimer(key string) Claimer {
	return &claimer{signKey: key}
}

func (c *claimer) IsValid(ctx context.Context, rawToken []byte) ([]byte, bool) {
	token := c.token(rawToken)
	if token == nil {
		return nil, false
	}

	mapVal, _ := token.AsMap(ctx)
	mapBytes, _ := json.Marshal(mapVal)

	return mapBytes, jwt.Validate(token) == nil
}

func (c *claimer) GetModel(ctx context.Context) *Model {
	return ctx.Value(utils.AuthCtx).(*Model)
}

func (c *claimer) token(rawToken []byte) jwt.Token {
	tok, err := jwt.Parse(rawToken, jwt.WithKey(jwa.HS256, []byte(c.signKey)))
	if err != nil {
		return nil
	}

	return tok
}
