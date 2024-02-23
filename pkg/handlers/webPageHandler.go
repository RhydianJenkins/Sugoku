package handlers

import (
	"net/http"
	"text/template"

	"github.com/rhydianjenkins/sugoku/pkg/board"
)

type PageData struct {
	Title string
	Tiles [board.BoardSize][board.BoardSize]board.Tile
}

func WebPageHandler(writer http.ResponseWriter, request *http.Request) {
	b := board.NewEmptyBoard()

	template, parseErr := template.ParseFiles("pkg/public/templates/index.html")
	if parseErr != nil {
		panic(parseErr)
	}

	err := template.Execute(writer, PageData{
		Title: "Sugoku - A Sudoku solver in Go!",
		Tiles: b.GetTiles(),
	})

	if err != nil {
		panic(err)
	}
}
