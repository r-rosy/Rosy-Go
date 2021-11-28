package web

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"UserSystemOfRosy3.0/database"
)

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	tab := database.SqlInitInfo()
	rd := tab.CheckDetails(username, "name", "name")
	if len(rd) == 0 {
		w.Write([]byte("该用户不存在"))
	} else {
		rd = tab.CheckDetails(username, "name", "pass")
		val, _ := rd[0]["pass"].([]byte)
		pas := string(val)
		if pas == password {
			rand.Seed(time.Now().UnixNano())
			value := strconv.Itoa(rand.Intn(1000))
			expir := time.Now().AddDate(0, 0, 1)
			cookie := &http.Cookie{
				Name:    username,
				Value:   value,
				Expires: expir,
			}
			http.SetCookie(w, cookie)
			expire := fmt.Sprintf("%v", expir)
			cook := map[string]string{"name": username, "value": value, "expire": expire}
			tab = database.SqlInitCookie()
			tab.Insert(cook)
			w.Write([]byte("登录成功!"))
		} else {
			w.Write([]byte("密码错误，登录失败！"))
		}
	}
}
