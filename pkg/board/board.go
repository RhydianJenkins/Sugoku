package board

import "fmt"

const BoardSize int = 2
const NumValues int = BoardSize * BoardSize

type Board struct {
	tiles [BoardSize][BoardSize]Tile
}

func New() Board {
	tiles := [BoardSize][BoardSize]Tile{}

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			tiles[i][j] = NewTile(i, j)
		}
	}

	// TODO find a better way of initialising the board
	tiles[0][0].value = 1

	return Board{
		tiles,
	}
}

func (board Board) String() string {
	return fmt.Sprintf("%v", board.tiles)
}

func (board Board) GetTile(x, y int) Tile {
	return board.tiles[x][y]
}

func (board Board) GetTiles() [BoardSize][BoardSize]Tile {
	return board.tiles
}

func (board Board) GetRow(x int) [BoardSize]Tile {
	return board.tiles[x]
}

func (board Board) GetCol(y int) [BoardSize]Tile {
	col := [BoardSize]Tile{}

	for i := 0; i < BoardSize; i++ {
		col[i] = board.GetTile(i, y)
	}

	return col
}
