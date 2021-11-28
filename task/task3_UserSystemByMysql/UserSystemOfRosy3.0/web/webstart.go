package web

import (
	"net/http"
)

func WebStart() {
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
