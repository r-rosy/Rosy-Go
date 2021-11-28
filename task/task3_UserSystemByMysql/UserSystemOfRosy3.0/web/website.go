package web

import (
	"fmt"
	"net/http"

	"UserSystemOfRosy3.0/database"
)

func website(w http.ResponseWriter, r *http.Request) {

	if database.CheckCookie(r) {
		fmt.Fprintf(w, "针不戳，终于进来了\n我的UID是184951275\n上原神来打我呀！")
	} else {
		w.Write([]byte("请先实行正确的登录，孩纸"))
	}

}
