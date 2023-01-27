package pg_migration

import (
	"konntent-authentication-service/pkg/pg"
)

func Migrate(pgi pg.Instance, models ...interface{}) error {
	for _, model := range models {
		err := pgi.Open().Model(model).CreateTable(model.(pg.Model).Opts())
		if err != nil {
			return err
		}
	}

	return nil
}
