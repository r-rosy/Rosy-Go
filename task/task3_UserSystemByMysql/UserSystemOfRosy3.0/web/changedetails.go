package web

import (
	"net/http"

	"UserSystemOfRosy3.0/database"
)

func ChangeDetails(w http.ResponseWriter, r *http.Request) {
	name := r.Cookies()[0].Name
	tab := database.SqlInitInfo()
	newpassword := r.FormValue("newpassword")
	newniname := r.FormValue("newniname")
	if database.CheckCookie(r) {
		if newpassword != "" {
			tab.ChangeWithConditon("pass", newpassword, "name", name)
		}
		if newniname != "" {

			tab.ChangeWithConditon("niname", newniname, "name", name)
		}
		w.Write([]byte("修改成功"))
	} else {
		w.Write([]byte("请先登录"))
	}
}
