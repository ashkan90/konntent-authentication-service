package pg

import "github.com/go-pg/pg/v10/orm"

type Model interface {
	String() string
	Opts() *orm.CreateTableOptions
}
