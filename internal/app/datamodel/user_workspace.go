package datamodel

import "github.com/go-pg/pg/v10/orm"

type UserWorkspace struct {
	UserID      int
	WorkspaceID int
}

func NewUserWorkspaceModel(uid, wid int) UserWorkspace {
	return UserWorkspace{
		UserID:      uid,
		WorkspaceID: wid,
	}
}

func (uw *UserWorkspace) String() string {
	return "UserWorkspace"
}

func (uw *UserWorkspace) Opts() *orm.CreateTableOptions {
	return &orm.CreateTableOptions{Temp: false, IfNotExists: true}
}

func (uw *UserWorkspace) ToWorkspace() *Workspace {
	return &Workspace{
		ID:     uw.WorkspaceID,
		UserID: uw.UserID,
	}
}
