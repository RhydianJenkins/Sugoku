package board

import (
	"fmt"
	"math/rand"
	"strings"
)

const BoardSize int = 9
const BlockSize int = 3

type Board struct {
	tiles                [BoardSize][BoardSize]Tile
	numPrePopulatedTiles int
	history              HistoryStack
}

func NewEmptyBoard() Board {
	tileValues := [BoardSize][BoardSize]int{}
	return NewBoard(tileValues)
}

func NewBoard(tileValues [BoardSize][BoardSize]int) Board {
	tiles := [BoardSize][BoardSize]Tile{}
	numPrePopulatedTiles := 0
	// Begin with an empty history, assuming the given board is solvable
	history := HistoryStack{}

	for x := range BoardSize {
		for y := range BoardSize {
			newTile := NewTile(x, y)
			newTile.Value = tileValues[x][y]
			tiles[x][y] = newTile

			if newTile.Value != Empty {
				numPrePopulatedTiles++
			}
		}
	}

	return Board{
		tiles,
		numPrePopulatedTiles,
		history,
	}
}

func (board *Board) Solve(numIterations int) error {
	for range numIterations {
		allTilesPopulated, _ := board.allTilesPopulated()
		isValid, _ := board.isValid()
		if allTilesPopulated && isValid {
			return nil
		}

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
	}

	return nil
}

func (board Board) allTilesPopulated() (solved bool, message string) {
	for x := range BoardSize {
		for y := range BoardSize {
			tile := board.GetTile(x, y)
			if tile.isEmpty() {
				return false, fmt.Sprintf("Tile at x:%d, y:%d is empty", x, y)
			}
		}
	}

	return true, "Board is solved"
}

func (board Board) isValid() (isValid bool, message string) {
	// check columns are unique
	for x := range BoardSize {
		col := board.GetCol(x)
		seenColValues := make(map[int]bool)

		for y := range BoardSize {
			val := col[y].Value
			if val == Empty {
				continue
			}

			if seenColValues[val] {
				return false, fmt.Sprintf("Val %d (x:%d, y:%d) is not unique in col", val, x, y)
			}

			seenColValues[val] = true
		}
	}

	// check rows are unique
	for y := range BoardSize {
		row := board.GetRow(y)
		seenRowValues := make(map[int]bool)

		for x := range BoardSize {
			val := row[x].Value
			if val == Empty {
				continue
			}

			if seenRowValues[val] {
				return false, fmt.Sprintf("Val %d (x:%d, y:%d) is not unique in row", val, x, y)
			}

			seenRowValues[val] = true
		}
	}

	// check blocks are unique
	seenBlockValues := make(map[int]bool)
	numBlocks := BoardSize / BlockSize
	for i := range numBlocks {
		for j := range numBlocks {
			block := board.GetBlock(i, j)
			for _, tile := range block {
				if tile.isEmpty() {
					continue
				}

				if tile.Value > BoardSize {
					return false, fmt.Sprintf("Val %d (x:%d, y:%d) is greater than board size", tile.Value, tile.X, tile.Y)
				}

				if tile.Value <= 0 {
					return false, fmt.Sprintf("Val %d (x:%d, y:%d) must be a positive integer", tile.Value, tile.X, tile.Y)
				}

				if seenBlockValues[tile.Value] {
					return false, fmt.Sprintf("Val %d (x:%d, y:%d) is not unique in block", tile.Value, tile.X, tile.Y)
				}
			}
		}
	}

	return true, "Board is valid"
}

func (board *Board) findLowestEntropyTiles() []*Tile {
	lowestEntropy := BoardSize
	lowestEntropyTiles := []*Tile{}

	for x := range BoardSize {
		for y := range BoardSize {
			tile := board.GetTile(x, y)
			tile.possibleValues = calculatePossibleValues(*board, x, y)
			tileEntropy := tile.GetEntropy()

			if tileEntropy < lowestEntropy && tileEntropy > 0 {
				lowestEntropy = tileEntropy
			}
		}
	}

	for x := range BoardSize {
		for y := range BoardSize {
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

	board.SetTileValue(randomTile.X, randomTile.Y, randomValue)

	return nil
}

func (board Board) String() string {
	var sb strings.Builder
	tiles := board.GetTiles()

	for x := range BoardSize {
		for y := range BoardSize {
			sb.WriteString(fmt.Sprintf("[%d]", tiles[x][y].Value))
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func (board *Board) GetBlock(x, y int) []*Tile {
	block := []*Tile{}

	for i := range BlockSize {
		for j := range BlockSize {
			block = append(block, board.GetTile(x*BlockSize+i, y*BlockSize+j))
		}
	}

	return block
}

func (board *Board) SetTileValue(x, y, val int) {
	board.history.push(board.GetTile(x, y))
	board.tiles[x][y].Value = val
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

	for i := range BoardSize {
		col[i] = *board.GetTile(i, y)
	}

	return col
}

func (board Board) GetHistory() []*Tile {
	return board.history.tiles
}
