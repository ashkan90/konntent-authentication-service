package datamodel

import (
	"github.com/go-pg/pg/v10/orm"
	"time"
)

type User struct {
	ID         int          `pg:",fk"`
	Email      string       `pg:",notnull"`
	Password   string       `pg:",notnull"`
	DeletedAt  time.Time    `pg:",soft_delete"`
	Account    *UserAccount `pg:"rel:has-one,fk:id"`
	Workspaces []*Workspace `pg:"many2many:user_workspaces"`
}

func (u *User) String() string {
	return "User"
}

func (u *User) Opts() *orm.CreateTableOptions {
	return &orm.CreateTableOptions{Temp: false, IfNotExists: true}
}
