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

func (board Board) findLowestEntropyTiles() []Tile {
	lowestEntropy := NumValues
	lowestEntropyTiles := []Tile{}

	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			tile := board.GetTile(x, y)
			possibleValues := calculatePossibleValues(board, x, y)
			entropy := len(possibleValues)

			if entropy < lowestEntropy && entropy > 0 {
				lowestEntropy = entropy
			}

			tile.entropy = entropy
			tile.possibleValues = possibleValues

			fmt.Println("board.GetTile(0,1) in first loop:", board.GetTile(0, 1).entropy)
			fmt.Println("tile.enropy in first loop:", tile.entropy)
		}
	}

	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			tile := board.GetTile(x, y)

			if tile.entropy == lowestEntropy {
				lowestEntropyTiles = append(lowestEntropyTiles, tile)
			}
		}
	}

	return lowestEntropyTiles
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
