package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/rhydianjenkins/sugoku/pkg/handlers"
)

func getPort() *string {
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	return port
}

func main() {
	port := getPort()

	fs := http.FileServer(http.Dir("pkg/public/styles"))
	http.Handle("/public/styles/", http.StripPrefix("/public/styles/", fs))

	http.HandleFunc("/api/test", handlers.HelloWorldHandler)
	http.HandleFunc("/", handlers.WebPageHandler)

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server listening on port " + *port)
}
