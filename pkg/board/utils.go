package board

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
