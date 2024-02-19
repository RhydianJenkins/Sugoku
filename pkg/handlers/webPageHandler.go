package handlers

import (
	"net/http"
	"text/template"
)

type PageData struct {
	Title string
}

func WebPageHandler(writer http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("pkg/public/templates/index.html")

	if err != nil {
		panic(err)
	}

	err = template.Execute(writer, PageData{
		Title: "Sugoku - A Sudoku solver in Go!",
	})

	if err != nil {
		panic(err)
	}
}
