package datamodel

import "github.com/go-pg/pg/v10/orm"

type UserAccount struct {
	ID             int `pg:",pk"`
	UserID         int `pg:",fk:id"`
	FirstName      string
	LastName       string
	JobDescription string
}

func (ua *UserAccount) String() string {
	return "UserAccount"
}

func (ua *UserAccount) Opts() *orm.CreateTableOptions {
	return &orm.CreateTableOptions{Temp: false, IfNotExists: true}
}
