package board

import "fmt"

const Empty int = 0

type TileVal struct {
	x, y, val int
}

type Tile struct {
	x              int
	y              int
	value          int
	possibleValues []int
}

func NewTile(x, y int) Tile {
	return Tile{
		x:              x,
		y:              y,
		value:          Empty,
		possibleValues: []int{},
	}
}

func (tile Tile) isEmpty() bool {
	return tile.value == Empty
}

func (tile Tile) String() string {
	return fmt.Sprintf("[%d]", tile.value)
}

func (tile Tile) GetEntropy() int {
	return len(tile.possibleValues)
}
