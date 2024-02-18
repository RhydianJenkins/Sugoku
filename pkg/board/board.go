package board

import "fmt"

const BoardSize int = 3
const BlockSize int = 3
const NumValues int = BoardSize * BlockSize

type Board struct {
	tiles                [BoardSize][BoardSize]Tile
	numPrePopulatedTiles int
}

func NewBoard(tileValues []TileVal) Board {
	tiles := [BoardSize][BoardSize]Tile{}

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			tiles[i][j] = NewTile(i, j)
		}
	}

	for _, tileVal := range tileValues {
		tiles[tileVal.x][tileVal.y].value = tileVal.val
	}

	return Board{
		tiles:                tiles,
		numPrePopulatedTiles: len(tileValues),
	}
}

func (board *Board) findLowestEntropyTiles() []*Tile {
	lowestEntropy := NumValues
	lowestEntropyTiles := []*Tile{}

	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			tile := board.GetTile(x, y)
			tile.possibleValues = calculatePossibleValues(*board, x, y)
			tileEntropy := tile.GetEntropy()

			if tileEntropy < lowestEntropy && tileEntropy > 0 {
				lowestEntropy = tileEntropy
			}
		}
	}

	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			tile := board.GetTile(x, y)

			if tile.GetEntropy() == lowestEntropy {
				lowestEntropyTiles = append(lowestEntropyTiles, tile)
			}
		}
	}

	return lowestEntropyTiles
}

func (board Board) String() string {
	return fmt.Sprintf("%v", board.tiles)
}

func (board *Board) GetBlock(x, y int) []*Tile {
	block := []*Tile{}
	return block
}

func (board *Board) GetTile(x, y int) *Tile {
	return &board.tiles[x][y]
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
		col[i] = *board.GetTile(i, y)
	}

	return col
}
