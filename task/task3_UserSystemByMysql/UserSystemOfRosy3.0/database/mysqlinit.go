package database

import (
	"github.com/typenameman/gotools/mysql"
)

func SqlInitInfo() *mysql.Table {
	db := mysql.NewDB(Url, "user")
	tab := &mysql.Table{
		Name:  "info",
		Types: Types1,
		Sqldb: db,
	}
	return tab
}
func SqlInitCookie() *mysql.Table {
	db := mysql.NewDB(Url, "user")
	tab := &mysql.Table{
		Name:  "cookie",
		Types: Types2,
		Sqldb: db,
	}
	return tab
}
