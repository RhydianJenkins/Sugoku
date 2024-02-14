package main

import (
	"net/http"

	"github.com/rhydianjenkins/sugoku/pkg/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("/pkg/public"))
	http.Handle("/public", http.StripPrefix("/public", fs))

	http.HandleFunc("/api/test", handlers.HelloWorldHandler)
	http.HandleFunc("/", handlers.WebPageHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	println("Server listening on port 8080...")
}
