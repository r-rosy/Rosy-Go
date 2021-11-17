package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	username string
	password string
	niname   string
	cookie   http.Cookie
}

var minimysql = make(map[string]*User)

func website(w http.ResponseWriter, r *http.Request) {
	for a, b := range minimysql {
		cookie, err := r.Cookie(a)
		if err == nil && cookie.Value == b.cookie.Value {
			fmt.Fprintf(w, "亲爱的用户%s:\n针不戳，终于进来了\n我的UID是184951275\n上原神来打我呀！", b.niname)
			goto THIS
		}
	}
	w.Write([]byte("请先实行正确的登录，孩纸"))
THIS:
}
func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if a, b := minimysql[username]; b {
		if a.password == password {
			rand.Seed(time.Now().UnixNano())
			value := strconv.Itoa(rand.Intn(1000))
			expire := time.Now().AddDate(0, 0, 1)
			cookie := &http.Cookie{
				Name:    username,
				Value:   value,
				Expires: expire,
			}
			http.SetCookie(w, cookie)
			a.cookie = *cookie
			w.Write([]byte("登录成功，请进入“/website”"))
		} else {
			w.Write([]byte("密码错误，登陆失败"))
		}
	} else {
		w.Write([]byte("该用户不存在"))
	}
}
func register(w http.ResponseWriter, r *http.Request) {
	var use User
	user := &use
	username := r.FormValue("username")
	password := r.FormValue("password")
	niname := r.FormValue("niname")
	if username != "" && password != "" {
		user.username = username
		user.password = password
		user.niname = niname
		minimysql[username] = user
		w.Write([]byte("您已注册成功,请跳转至登录页面"))
	} else {
		w.Write([]byte("您未注册成功，并请保证您已同时输入正确的用户名和密码"))
	}
}
func ChangeDetails(w http.ResponseWriter, r *http.Request) {
	var username string
	bools := false
	for a, b := range minimysql {
		cookie, err := r.Cookie(a)
		if err == nil && cookie.Value == b.cookie.Value {
			username = cookie.Name
			bools = true
			goto THISS
		}
	}
	w.Write([]byte("请先登录"))
THISS:
	if bools {
		newname := r.FormValue("newname")
		newpassword := r.FormValue("newpassword")
		newniname := r.FormValue("newniname")
		user := minimysql[username]
		if newname != "" {
			user.username = newname
		}
		if newpassword != "" {
			user.password = newpassword
		}
		if newniname != "" {
			user.niname = newniname
		}
		w.Write([]byte("修改成功"))
	}
}
func WatchDetails(w http.ResponseWriter, r *http.Request) {
	var username string
	bools := false
	for a, b := range minimysql {
		cookie, err := r.Cookie(a)
		if (err == nil) && (cookie.Value == b.cookie.Value) {
			username = cookie.Name
			bools = true
			goto THIS
		}
	}
	w.Write([]byte("请进入登陆页面正确登录"))
THIS:
	if bools {
		user := minimysql[username]
		fmt.Fprintf(w, "登录成功请查看信息:\n用户名：%s\n昵称：%s", user.username, user.niname)
	}
}
func main() {
	servemux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":211",
		Handler: servemux,
	}
	servemux.HandleFunc("/website", website)
	servemux.HandleFunc("/login", login)
	servemux.HandleFunc("/register", register)
	servemux.HandleFunc("/WatchDetails", WatchDetails)
	servemux.HandleFunc("/ChangeDetails", ChangeDetails)
	server.ListenAndServe()

}
