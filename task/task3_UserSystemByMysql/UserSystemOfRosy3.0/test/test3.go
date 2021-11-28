package test

import (
	"fmt"

	"UserSystemOfRosy3.0/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/typenameman/gotools/handle"
)

func Test3() {
	tab := database.SqlInitCookie()
	res := tab.CheckDetails("muxi", "name", "value")
	s := handle.InterfaceToString(res[0]["value"])
	fmt.Println(s)
}
