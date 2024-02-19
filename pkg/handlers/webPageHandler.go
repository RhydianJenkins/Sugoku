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
	template, err := template.ParseFiles("pkg/public/templates/index.html")

	if err != nil {
		panic(err)
	}

	b := board.NewBoard([]board.TileVal{})

	for i := 0; i < board.BoardSize*board.BoardSize; i++ {
		b.SolveOneStep()
	}

	err = template.Execute(writer, PageData{
		Title: "Sugoku - A Sudoku solver in Go!",
		Tiles: b.GetTiles(),
	})

	if err != nil {
		panic(err)
	}
}
