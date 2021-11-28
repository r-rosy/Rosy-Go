package web

import (
	"fmt"
	"net/http"

	"UserSystemOfRosy3.0/database"
)

func WatchDetails(w http.ResponseWriter, r *http.Request) {
	tab := database.SqlInitInfo()
	na := r.Cookies()[0].Name
	res := tab.CheckDetails(na, "name", "name", "niname")
	if database.CheckCookie(r) {
		fmt.Fprintf(w, "登录成功请查看信息:\n用户名：%s\n昵称：%s", res[0]["name"], res[0]["niname"])
	} else {
		w.Write([]byte("请进入登陆页面正确登录"))
	}

}
