package board

import (
	"fmt"
)

func boardIsValid(board Board) (isValid bool, message string) {
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

func calculatePossibleValues(board Board, x, y int) []int {
	tile := board.tiles[x][y]

	if !tile.isEmpty() {
		return []int{}
	}

	row := board.GetRow(tile.X)
	col := board.GetCol(tile.Y)
	possibleValues := []int{}

	for i := 0; i <= BoardSize; i++ {
		possibleValues = append(possibleValues, i)
	}

	for _, val := range tile.BadValues {
		possibleValues[val] = Empty
	}

	for i := 0; i < BoardSize; i++ {
		if !row[i].isEmpty() {
			possibleValues[row[i].Value] = Empty
		}

		if !col[i].isEmpty() {
			possibleValues[col[i].Value] = Empty
		}
	}

	blockX := x / BlockSize
	blocky := y / BlockSize
	block := board.GetBlock(blockX, blocky)
	for _, t := range block {
		if t.Value != Empty {
			possibleValues[t.Value] = Empty
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
