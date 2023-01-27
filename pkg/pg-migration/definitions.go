package pg_migration

import (
	"konntent-authentication-service/internal/app/datamodel"
)

var MigrationModels = []interface{}{
	(*datamodel.User)(nil),
	(*datamodel.UserAccount)(nil),
	(*datamodel.Workspace)(nil),
	(*datamodel.UserWorkspace)(nil),
}
