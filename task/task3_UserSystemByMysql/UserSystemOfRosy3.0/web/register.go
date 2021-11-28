package web

import (
	"net/http"

	"UserSystemOfRosy3.0/database"
)

func register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	niname := r.FormValue("niname")
	if username != "" && password != "" && niname != "" {

		tab := database.SqlInitInfo()
		rd := tab.CheckDetails(username, "name", "name")
		if len(rd) != 0 {
			w.Write([]byte("该用户已存在"))
		} else {
			var Info = map[string]string{"name": username, "pass": password, "niname": niname}
			tab.Insert(Info)
			w.Write([]byte("您已注册成功,用户名作为您的身份凭证一经注册即不可修改，请跳转至登录页面进行登录"))
		}
	} else {
		w.Write([]byte("您未注册成功，并请保证您已同时输入正确的用户名和密码以及昵称"))
	}
}
