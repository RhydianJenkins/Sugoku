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

func (board Board) CalculatePossibleValues(x, y int) []int {
	tile := board.tiles[x][y]

	if tile.IsEmpty() {
		return []int{tile.value}
	}

	row := board.GetRow(tile.x)
	col := board.GetCol(tile.y)
	possibleValues := []int{}

	for i := 1; i <= NumValues; i++ {
		possibleValues = append(possibleValues, i)
	}

	fmt.Println("possibleValues: ", possibleValues)

	for i := 0; i < BoardSize; i++ {
		if !row[i].IsEmpty() {
			possibleValues[row[i].value-1] = Empty
		}

		if !col[i].IsEmpty() {
			possibleValues[col[i].value-1] = Empty
		}
	}

	return filterEmpty(possibleValues)
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

func (board Board) GetRow(i int) [BoardSize]Tile {
	return board.tiles[i]
}

func (board Board) GetCol(i int) [BoardSize]Tile {
	col := [BoardSize]Tile{}

	for j := 0; j < BoardSize; j++ {
		col[j] = board.GetTile(j, i)
	}

	return col
}

func filterEmpty(slice []int) []int {
	var filtered []int
	for _, value := range slice {
		if value != Empty {
			filtered = append(filtered, value)
		}
	}
	return filtered
}
