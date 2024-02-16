package board

import "fmt"

const Empty int = 0

type Tile struct {
	x     int
	y     int
	value int
}

func NewTile(x, y int) Tile {
	return Tile{
		x:     x,
		y:     y,
		value: Empty,
	}
}

func (tile Tile) isEmpty() bool {
	return tile.value == Empty
}

func (tile Tile) String() string {
	return fmt.Sprintf("[ %d ]", tile.value)
}
