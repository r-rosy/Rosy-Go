package main

import (
	_ "fmt"
	"html/template"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	tem, _ := template.ParseFiles("./ttt.html")
	tem.Execute(w, nil)
}
func main() {
	servemux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":211",
		Handler: servemux,
	}
	servemux.HandleFunc("/handle", handle)
	server.ListenAndServe()
}
