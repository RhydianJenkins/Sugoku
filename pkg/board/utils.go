package board

func solveOneStep(board Board) {
	panic("TODO: implement")

	// 1. get slice of lowest entropy tiles
	// 2. for each tile, get possible values
	// 3. if only one possible value, set the value
	// 4. if no possible value, return error
	// 5. select random tile from lowest entropy tiles
	// 6. select random value from possible values
	// 7. set the value
	// 8. ???
	// 9. profit
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
