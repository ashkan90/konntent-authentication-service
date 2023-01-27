package tokenizer

import (
	"github.com/golang-jwt/jwt/v4"
	configs "konntent-authentication-service/configs/app"
	"log"
	"time"
)

type ITokenizer interface {
	Tokenize(claims Claims) IToken
}

type IToken interface {
	Token() *Token
	Raw() string

	IsValid() bool
	IsExpired() bool
}

type Token struct {
	TokenRaw  string    `json:"token"`
	Valid     error     `json:"-"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type tokenizer struct {
	conf configs.JWTConfig
}

func NewTokenizer(cfg configs.JWTConfig) ITokenizer {
	return &tokenizer{conf: cfg}
}

func (t *tokenizer) Tokenize(claims Claims) IToken {
	return newToken(t.conf, claims)
}

func (t *Token) Token() *Token {
	return t
}

func (t *Token) Raw() string {
	return t.TokenRaw
}

func (t *Token) IsValid() bool {
	return t.Valid == nil
}

func (t *Token) IsExpired() bool {
	return t.ExpiresAt.Sub(time.Now().UTC()).Seconds() < 0
}

func newToken(cfg configs.JWTConfig, claims Claims) *Token {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(cfg.TTL) * time.Second))

	var tok = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var raw, err = tok.SignedString([]byte(cfg.SignKey))

	if err != nil {
		log.Println("newToken error: ", err.Error())
		return nil
	}

	return &Token{
		TokenRaw:  raw,
		Valid:     tok.Claims.Valid(),
		ExpiresAt: claims.ExpiresAt.Time,
	}
}
