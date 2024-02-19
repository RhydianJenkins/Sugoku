package board

import "fmt"

const Empty int = 0

type TileVal struct {
	x, y, val int
}

type Tile struct {
	X              int   `json:"x"`
	Y              int   `json:"y"`
	Value          int   `json:"value"`
	PossibleValues []int `json:"possibleValues"`
}

func NewTile(x, y int) Tile {
	return Tile{
		X:              x,
		Y:              y,
		Value:          Empty,
		PossibleValues: []int{},
	}
}

func (tile Tile) isEmpty() bool {
	return tile.Value == Empty
}

func (tile Tile) String() string {
	if tile.Value == Empty {
		return ""
	}

	return fmt.Sprintf("%d", tile.Value)
}

func (tile Tile) GetEntropy() int {
	return len(tile.PossibleValues)
}
