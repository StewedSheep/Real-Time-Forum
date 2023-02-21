package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {

	StartDb()

	server := &http.Server{
		Addr:    ":" + "8000",
		Handler: Handler(),
	}

	fmt.Println("Server running 8000")
	log.Fatal(server.ListenAndServe())
	fmt.Println("quit!")

}

func Handler() http.Handler {

	handler := http.FileServer(http.Dir("frontend/dist"))
	wsServer := NewWebsocketServer()
	go wsServer.Run()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_path := r.URL.Path

		//api calls
		if strings.HasPrefix(_path, "/api/v1/") {
			if strings.HasSuffix(_path, "login") {
				LoginAuth(w, r)
			}
			if strings.HasSuffix(_path, "logout") {
				ResetCookie(w, r)
			}
			if strings.HasSuffix(_path, "register") {
				ValidateRegister(w, r)
			}
			if strings.HasSuffix(_path, "totThread") {
				GetRandomLaw(w, r)
			}
			if strings.HasSuffix(_path, "threaddat") {
				CreateThread(w, r)
			}
			if strings.HasSuffix(_path, "ws") {
				ServeWs(wsServer, w, r)
			}
			return
		}

		// static files
		if strings.Contains(_path, ".") || _path == "/" {
			handler.ServeHTTP(w, r)
			return
		}

		// all 404 going to be served as root
		http.ServeFile(w, r, ("frontend/dist/index.html"))
	})
}
