package test

import (
	"fmt"

	"UserSystemOfRosy3.0/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/typenameman/gotools/handle"
)

func Test2() {
	tab := database.SqlInitInfo()
	res := tab.ShowAll()
	s := handle.InterfaceToString(res[0]["name"])
	fmt.Println(s)
}
