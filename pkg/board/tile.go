package board

import "fmt"

const Empty int = 0
const InitialEntropy int = -1

type Tile struct {
	x              int
	y              int
	value          int
	entropy        int
	possibleValues []int
}

func NewTile(x, y int) Tile {
	return Tile{
		x:              x,
		y:              y,
		value:          Empty,
		entropy:        InitialEntropy,
		possibleValues: []int{},
	}
}

func (tile Tile) isEmpty() bool {
	return tile.value == Empty
}

func (tile Tile) String() string {
	return fmt.Sprintf("[ %d ]", tile.value)
}
