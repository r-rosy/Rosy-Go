package test

import (
	"UserSystemOfRosy3.0/database"
	_ "github.com/go-sql-driver/mysql"
)

func Test4() {
	tab := database.SqlInitInfo()

	tab.ChangeWithConditon("niname", "hop", "name", "muxi")
}
