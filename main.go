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
	go Echo()

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

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_path := r.URL.Path

		//api calls
		if strings.HasPrefix(_path, "/api/v1/") {
			if strings.HasSuffix(_path, "login") {
				fmt.Println("LoginAuth")
				LoginAuth(w, r)
			}
			if strings.HasSuffix(_path, "logout") {
				fmt.Println("ResetCookie")
				ResetCookie(w, r)
			}
			if strings.HasSuffix(_path, "register") {
				fmt.Println("ValidateRegister")
				ValidateRegister(w, r)
			}
			if strings.HasSuffix(_path, "totThread") {
				fmt.Println("GetAllThreads")
				GetAllThreads(w, r)
			}
			if strings.HasSuffix(_path, "totUsers") {
				fmt.Println("GetAllUsers")
				GetAllUsers(w, r)
			}
			if strings.HasSuffix(_path, "threaddat") {
				fmt.Println("CreateThread")
				CreateThread(w, r)
			}
			if strings.HasSuffix(_path, "commentdat") {
				fmt.Println("CreateComment")
				CreateComment(w, r)
			}
			if strings.HasSuffix(_path, "thread") {
				fmt.Println("GetThread")
				GetThread(w, r)
			}
			if strings.HasSuffix(_path, "ws") {
				fmt.Println("ServeWebSocket")
				WsHandler(w, r)
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
