package test

import (
	"fmt"

	"UserSystemOfRosy3.0/database"
	_ "github.com/go-sql-driver/mysql"
)

func Test1() {
	/*db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user")
	if err != nil {
		panic(err)
	}
	db.Exec("create table user;")
	*/
	tab := database.SqlInitInfo()
	/*var de = map[string]string{"name": "'123'", "pass": "'1'", "niname": "'1'", "cookie": "'12'"}
	tab.Insert(de)*/
	rd := tab.CheckDetails("123", "name", "*")
	fmt.Println(rd)
}
