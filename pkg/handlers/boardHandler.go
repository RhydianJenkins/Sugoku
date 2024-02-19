package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rhydianjenkins/sugoku/pkg/board"
)

type Response struct {
	Tiles [board.BoardSize][board.BoardSize]board.Tile `json:"tiles"`
}

// TODO /api/solve/1 solves just one step of a Board and /api/solve solves the whole board
func BoardHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	initialTileVals := []board.TileVal{}
	b := board.NewBoard(initialTileVals)

	for i := 0; i < board.BoardSize*board.BoardSize; i++ {
		b.SolveOneStep()
	}

	response := Response{
		Tiles: b.GetTiles(),
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
