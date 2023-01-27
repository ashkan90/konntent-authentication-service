package pg_rel_registration

import (
	"github.com/go-pg/pg/v10/orm"
	"konntent-authentication-service/internal/app/datamodel"
)

func Register() {
	orm.RegisterTable((*datamodel.UserWorkspace)(nil))
}
