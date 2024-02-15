package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/rhydianjenkins/sugoku/pkg/handlers"
)

func getPort() (port *int) {
	p := flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

	return p
}

func main() {
	port := getPort()

	fs := http.FileServer(http.Dir("pkg/public/styles"))
	http.Handle("/public/styles/", http.StripPrefix("/public/styles/", fs))

	http.HandleFunc("/api/solve", handlers.BoardHandler)
	http.HandleFunc("/", handlers.WebPageHandler)

	fmt.Println(fmt.Sprintf("Starting server on localhost:%d", *port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		panic(err)
	}
}
