package main

import (
	"fmt"
	"net/http"
)

type User struct {
	username string
	password string
	niname   string
}

var minimysql []*User

func Enter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WebsitesList:\n/register:注册\n/login:登录\n/ChangeDetails:修改个人信息\n/WatchDetails:查看个人信息"))
}
func website(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("终于进来了吧，上原神来打我呀！！！"))
}
func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	n := 0
	for _, a := range minimysql {
		if a.username == username {
			n++
			if a.password == password {
				w.Write([]byte("登录成功，请进入“/website”"))
			} else {
				w.Write([]byte("登录失败，密码错误"))
			}
		}
	}
	if n == 0 {
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
		minimysql = append(minimysql, user)
		w.Write([]byte("您已注册成功,请跳转至登录页面"))
	} else {
		w.Write([]byte("您正在注册，并请保证您已同时输入用户名和密码"))
	}
}
func ChangeDetails(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	newname := r.FormValue("newname")
	newpassword := r.FormValue("newpassword")
	newniname := r.FormValue("newniname")
	n := 0
	for _, a := range minimysql {
		if a.username == username {
			n++
			if a.password == password {
				if newname != "" {
					a.username = newname
				}
				if newpassword != "" {
					a.password = newpassword
				}
				if newniname != "" {
					a.niname = newniname
				}
				w.Write([]byte("修改成功"))
			} else {
				w.Write([]byte("修改失败，原密码错误"))
			}
		}
	}
	if n == 0 {
		w.Write([]byte("该用户不存在"))
	}
}
func WatchDetails(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	n := 0
	for _, a := range minimysql {
		if a.username == username {
			n++
			if a.password == password {
				fmt.Fprintf(w, "登录成功请查看信息:\n用户名：%s\n昵称：%s", a.username, a.niname)

			} else {
				w.Write([]byte("登录失败，密码错误"))
			}
		}
	}
	if n == 0 {
		w.Write([]byte("该用户不存在"))
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
	servemux.HandleFunc("/Enter", Enter)
	servemux.HandleFunc("/WatchDetails", WatchDetails)
	servemux.HandleFunc("/ChangeDetails", ChangeDetails)
	server.ListenAndServe()

}
