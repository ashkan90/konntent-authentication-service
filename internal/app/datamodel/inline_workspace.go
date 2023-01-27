package datamodel

import "github.com/go-pg/pg/v10/orm"

type Workspace struct {
	ID     int
	UserID int
}

func NewWorkspaceModel(uid int) Workspace {
	return Workspace{UserID: uid}
}

func (w *Workspace) String() string {
	return "Workspace"
}

func (w *Workspace) Opts() *orm.CreateTableOptions {
	return &orm.CreateTableOptions{Temp: false, IfNotExists: true}
}
