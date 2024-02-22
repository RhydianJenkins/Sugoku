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
		tiles[tileVal.x][tileVal.y].Value = tileVal.val
	}

	return Board{
		tiles:                tiles,
		numPrePopulatedTiles: len(tileValues),
	}
}

func (board *Board) Solve() (err bool, msg string) {
	for {
		err, _ := board.solveOneStep()
		if err {
			poppedTile, empty := board.history.pop()
			if empty {
				return true, "Board is unsolvable"
			}

			badValue := poppedTile.Value
			poppedTile.Value = Empty
			poppedTile.BadValues = append(poppedTile.BadValues, badValue)
		}

		if board.isSolved() {
			return false, "Board is solved"
		}
	}
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

func (board *Board) solveOneStep() (error bool, msg string) {
	lowestEntropyTiles := board.findLowestEntropyTiles()

	if len(lowestEntropyTiles) == 0 {
		return true, "No tiles found with lowest entropy"
	}

	randomTileIndex := rand.Intn(len(lowestEntropyTiles))
	randomTile := lowestEntropyTiles[randomTileIndex]
	randomValueIndex := rand.Intn(len(randomTile.possibleValues))
	randomValue := randomTile.possibleValues[randomValueIndex]

	randomTile.Value = randomValue

	board.history.push(randomTile)

	return false, fmt.Sprintf("Tile (%d, %d) given value %d", randomTile.X, randomTile.Y, randomTile.Value)
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
