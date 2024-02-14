package main

import (
	"encoding/json"
	"net/http"

	"github.com/rhydianjenkins/sugoku/server"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var result = server.TestFn()
	print(result)

	response := Response{
		Message: "Hello, World!",
		Status:  200,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)

	println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
