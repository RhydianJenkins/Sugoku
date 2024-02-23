package board

import (
	"fmt"
	"math/rand"
)

const BoardSize int = 9
const BlockSize int = 3

type Board struct {
	tiles                [BoardSize][BoardSize]Tile
	numPrePopulatedTiles int
	history              HistoryStack
}

func NewBoard(tileValues []TileVal) Board {
	tiles := [BoardSize][BoardSize]Tile{}

	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			tiles[x][y] = NewTile(x, y)
		}
	}

	for _, tileVal := range tileValues {
		tiles[tileVal.X][tileVal.Y].Value = tileVal.Val
	}

	return Board{
		tiles:                tiles,
		numPrePopulatedTiles: len(tileValues),
	}
}

func (board *Board) Solve(numIterations int) error {
	for i := 0; i < numIterations; i++ {
		err := board.solveOneStep()

		if err != nil {
			poppedTile := board.history.pop()
			if poppedTile == nil {
				return fmt.Errorf("Tried to backtrack with empty history. Board is unsolvable")
			}

			badValue := poppedTile.Value
			poppedTile.Value = Empty
			poppedTile.BadValues = append(poppedTile.BadValues, badValue)
		}

		if board.isSolved() {
			return nil
		}
	}

	return nil
}

func (board Board) isSolved() bool {
	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			tile := board.GetTile(x, y)
			if tile.isEmpty() {
				return false
			}
		}
	}

	return true
}

func (board Board) isValid() (isValid bool, message string) {
	seenValues := make(map[int]bool)

	// check columns are unique
	for x := 0; x < BoardSize; x++ {
		col := board.GetCol(x)
		seenColValues := make(map[int]bool)

		for y := 0; y < BoardSize; y++ {
			val := col[y].Value
			if val == Empty {
				continue
			}

			if seenColValues[val] {
				return false, fmt.Sprintf("Val %d is not unique in col", val)
			}

			if seenValues[val] {
				return false, fmt.Sprintf("Val %d is not unique in board", val)
			}

			seenColValues[val] = true
			seenValues[val] = true
		}
	}

	// check rows are unique
	for y := 0; y < BoardSize; y++ {
		row := board.GetRow(y)
		seenRowValues := make(map[int]bool)

		for x := 0; x < BoardSize; x++ {
			val := row[x].Value
			if val == Empty {
				continue
			}

			if seenRowValues[val] {
				return false, fmt.Sprintf("Val %d is not unique in row %d", val, x)
			}

			seenRowValues[val] = true
		}
	}

	return true, "Board is valid"
}

func (board *Board) findLowestEntropyTiles() []*Tile {
	lowestEntropy := BoardSize
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

func (board *Board) solveOneStep() error {
	lowestEntropyTiles := board.findLowestEntropyTiles()

	if len(lowestEntropyTiles) == 0 {
		return fmt.Errorf("No tiles found with lowest entropy")
	}

	randomTileIndex := rand.Intn(len(lowestEntropyTiles))
	randomTile := lowestEntropyTiles[randomTileIndex]
	randomValueIndex := rand.Intn(len(randomTile.possibleValues))
	randomValue := randomTile.possibleValues[randomValueIndex]

	randomTile.Value = randomValue

	board.history.push(randomTile)

	return nil
}

func (board Board) String() string {
	return fmt.Sprintf("%v", board.tiles)
}

func (board *Board) GetBlock(x, y int) []*Tile {
	block := []*Tile{}

	for i := 0; i < BlockSize; i++ {
		for j := 0; j < BlockSize; j++ {
			block = append(block, board.GetTile(x*BlockSize+i, y*BlockSize+j))
		}
	}

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

func (board Board) GetHistory() []*Tile {
	return board.history.tiles
}
