package board

import (
	"fmt"
	"math/rand"
)

func boardIsValid(board Board) (isValid bool, message string) {
	seenValues := make(map[int]bool)

	// check columns are unique
	for x := 0; x < BoardSize; x++ {
		col := board.GetCol(x)
		seenColValues := make(map[int]bool)

		for y := 0; y < BoardSize; y++ {
			val := col[y].value
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
			val := row[x].value
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

func solveOneStep(board *Board) {
	lowestEntropyTiles := board.findLowestEntropyTiles()

	if len(lowestEntropyTiles) == 0 {
		panic("No solution found. TODO backtrack")
	}

	randomTileIndex := rand.Intn(len(lowestEntropyTiles))
	randomTile := lowestEntropyTiles[randomTileIndex]
	randomValueIndex := rand.Intn(len(randomTile.possibleValues))
	randomValue := randomTile.possibleValues[randomValueIndex]

	randomTile.value = randomValue
}

func calculatePossibleValues(board Board, x, y int) []int {
	tile := board.tiles[x][y]

	if !tile.isEmpty() {
		return []int{}
	}

	row := board.GetRow(tile.x)
	col := board.GetCol(tile.y)
	possibleValues := []int{}

	for i := 1; i <= NumValues; i++ {
		possibleValues = append(possibleValues, i)
	}

	for i := 0; i < BoardSize; i++ {
		if !row[i].isEmpty() {
			possibleValues[row[i].value-1] = Empty
		}

		if !col[i].isEmpty() {
			possibleValues[col[i].value-1] = Empty
		}
	}

	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			val := board.tiles[x][y].value

			if val != Empty {
				possibleValues[val-1] = Empty
			}
		}
	}

	return filterEmpty(possibleValues)
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
