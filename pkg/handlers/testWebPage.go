package handlers

import (
	"net/http"
	"text/template"
)

type PageData struct {
	Title string
	Body  string
}

func WebPageHandler(writer http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("pkg/templates/testPage.html")

	if err != nil {
		panic(err)
	}

	err = template.Execute(writer, PageData{
		Title: "My Page Title",
		Body:  "This is the body of my web page.",
	})

	if err != nil {
		panic(err)
	}
}
