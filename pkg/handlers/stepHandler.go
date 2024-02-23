package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rhydianjenkins/sugoku/pkg/board"
)

type StepRequest struct {
	Tiles [board.BoardSize][board.BoardSize]int `json:"tiles"`
}

type StepResponse struct {
	Tiles [board.BoardSize][board.BoardSize]board.Tile `json:"tiles"`
}

func StepHandler(writer http.ResponseWriter, request *http.Request) {
	var stepRequest StepRequest

	err := json.NewDecoder(request.Body).Decode(&stepRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	b := board.NewBoard(stepRequest.Tiles)
	b.Solve(1)

	response := StepResponse{
		Tiles: b.GetTiles(),
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
