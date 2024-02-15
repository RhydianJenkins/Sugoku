package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rhydianjenkins/sugoku/pkg/board"
)

type Response struct {
	BoardStr string                                       `json:"boardString"`
	Tiles    [board.BoardSize][board.BoardSize]board.Tile `json:"tiles"`
}

func BoardHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	b := board.New()

	response := Response{
		BoardStr: b.String(),
		Tiles:    b.GetTiles(),
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
