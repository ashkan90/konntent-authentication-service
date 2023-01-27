package pg

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/extra/pgsegment/v10"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/pgjson"
	"go.uber.org/zap"
	"konntent-authentication-service/configs/app"
	"konntent-authentication-service/pkg/pg/hooks"
	"log"
)

type Instance interface {
	Open() *pg.DB
}

type instance struct {
	db *pg.DB
}

func init() {
	pgjson.SetProvider(pgsegment.NewJSONProvider())
}

func NewPGInstance(l *zap.Logger, conf app.PGSettings) (Instance, error) {
	var (
		url = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			conf.User, conf.Password, "localhost", 5433, "konntent-auth")
		opt, _ = pg.ParseURL(url)
	)

	log.Println(url)

	var i = &instance{
		db: pg.Connect(opt),
	}

	err := i.db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	if conf.Debug {
		i.db.AddQueryHook(hooks.NewDebugHook(l))
	}

	return i, nil
}

func (i *instance) Open() *pg.DB {
	return i.db
}
