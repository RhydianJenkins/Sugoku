package board

import (
	"reflect"
	"testing"
)

func createBoard() Board {
	b := NewEmptyBoard()
	b.SetTileValue(0, 0, 1)

	return b
}

func TestGetCol(t *testing.T) {
	colNum := BoardSize - 1
	b := createBoard()
	col := b.GetCol(colNum)
	tiles := b.GetTiles()

	for i := 0; i < BoardSize; i++ {
		if col[i].X != tiles[i][colNum].X || col[i].Y != tiles[i][colNum].Y {
			t.Errorf("Expected 0, got %v", col[i])
		}
	}
}

func TestCalculatePossibleValuesWithNonEmptyVal(t *testing.T) {
	returned := calculatePossibleValues(createBoard(), 0, 0)
	expected := []int{}

	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("Expected %v, got %v", expected, returned)
	}
}

func TestCalculatePossibleValuesWithBadValues(t *testing.T) {
	board := createBoard()
	board.GetTile(2, 3).BadValues = append(board.GetTile(2, 3).BadValues, 1)
	board.GetTile(2, 3).BadValues = append(board.GetTile(2, 3).BadValues, 4)
	returned := calculatePossibleValues(board, 2, 3)
	expected := []int{2, 3, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("Expected %v, got %v", expected, returned)
	}
}

func TestCalculatePossibleValuesWithEmptyVal(t *testing.T) {
	returned := calculatePossibleValues(createBoard(), 0, 1)
	expected := []int{2, 3, 4, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("Expected %v, got %v", expected, returned)
	}
}

func TestFilterEmpty(t *testing.T) {
	expected := []int{1, 2, 3, 5}
	returned := filterEmpty([]int{0, 1, 2, 3, 0, 5})

	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("Expected %v, got %v", expected, returned)
	}
}

func TestFindLowestEntropyTiles(t *testing.T) {
	board := NewEmptyBoard()
	board.SetTileValue(0, 0, 1)
	board.SetTileValue(1, 1, 2)
	board.SetTileValue(4, 4, 2)
	board.SetTileValue(5, 4, 5)
	numReturned := len(board.findLowestEntropyTiles())
	numExpected := 1

	if numReturned != numExpected {
		t.Errorf("Expected %v, but got %v", numExpected, numReturned)
	}

	lowest := board.findLowestEntropyTiles()[0]
	if lowest.X != 0 || lowest.Y != 4 {
		t.Errorf("Expected lowest tile to be (0, 4) but got (%v, %v)", lowest.X, lowest.Y)
	}

	expected := []int{3, 4, 6, 7, 8, 9}
	if !reflect.DeepEqual(expected, lowest.possibleValues) {
		t.Errorf("Expected possible values to be %v, but got %v", expected, lowest.possibleValues)
	}
}

func TestFindLowestEntropyTilesOnEmptyBoard(t *testing.T) {
	board := NewEmptyBoard()
	numReturned := len(board.findLowestEntropyTiles())
	numExpected := BoardSize * BoardSize

	if numReturned != numExpected {
		t.Errorf("Expected %v, got %v", numExpected, numReturned)
	}
}

func TestSolveOneStep(t *testing.T) {
	board := createBoard()
	for i := 0; i < BoardSize-board.numPrePopulatedTiles; i++ {
		err := board.solveOneStep()
		if err != nil {
			t.Errorf("Unable to solve board. Error: %v", err)
			break
		}
	}

	isValid, message := board.isValid()

	if !isValid {
		t.Errorf("Expected board to be valid, got invalid with message %v", message)
	}
}

func TestSolveOneAddsToHistory(t *testing.T) {
	board := createBoard()

	if board.history.isEmpty() {
		t.Errorf("Expected history to not be empty")
	}

	board.solveOneStep()
	if len(board.history.tiles) != 2 {
		t.Errorf("Expected history to have 2 tiles, got %v", len(board.history.tiles))
	}
}

func TestBoardIsValid(t *testing.T) {
	board := createBoard()
	isValid, message := board.isValid()
	if !isValid {
		t.Errorf("Expected starting board to be valid, got invalid with message %v", message)
	}

	board.GetTile(0, 1).Value = 1
	isValid, message = board.isValid()
	if isValid {
		t.Errorf("Expected modified board to be invalid, got valid with message %v", message)
	}
}

func TestGetBlock(t *testing.T) {
	board := createBoard()
	board.GetTile(3, 4).Value = 2
	block := board.GetBlock(1, 0)
	expected := []*Tile{
		board.GetTile(3, 0),
		board.GetTile(3, 1),
		board.GetTile(3, 2),
		board.GetTile(4, 0),
		board.GetTile(4, 1),
		board.GetTile(4, 2),
		board.GetTile(5, 0),
		board.GetTile(5, 1),
		board.GetTile(5, 2),
	}

	if !reflect.DeepEqual(expected, block) {
		t.Errorf("Expected %v, got %v", expected, block)
	}
}

func TestSolve(t *testing.T) {
	board := NewEmptyBoard()
	allTilesPopulatedBefore, _ := board.allTilesPopulated()

	if allTilesPopulatedBefore {
		t.Errorf("Expected board to be unsolved before solving")
	}

	err := board.Solve(999)
	isValid, validMessage := board.isValid()
	allTilesPopulated, solvedMessage := board.allTilesPopulated()

	if err != nil {
		t.Errorf("Solve returned error with message '%v'", err)
	}

	if !isValid {
		t.Errorf("Expected board to be valid, got invalid with message '%v'", validMessage)
	}

	if !allTilesPopulated {
		t.Errorf("board.isSolved() returned false with message %v", solvedMessage)
	}
}

func TestFullySolvedIsSolvedAndValid(t *testing.T) {
	board := NewEmptyBoard()

	for x := 0; x < BoardSize; x++ {
		offset := ((x * BlockSize) % BoardSize) + (x / BlockSize)
		for y := 0; y < BoardSize; y++ {
			val := (y+offset)%BoardSize + 1
			board.SetTileValue(x, y, val)
		}
	}

	allTilesPopulated, solvedMessage := board.allTilesPopulated()
	if !allTilesPopulated {
		t.Errorf("Board is not solved with message '%v'", solvedMessage)
	}

	isValid, message := board.isValid()
	if !isValid {
		t.Errorf("Expected board to be valid but got invalid with message '%v'", message)
	}
}

func TestPartiallySolvedWithBacktracking(t *testing.T) {
	board := NewEmptyBoard()

	for x := 0; x < BoardSize; x++ {
		offset := ((x * BlockSize) % BoardSize) + (x / BlockSize)
		for y := 0; y < BoardSize; y++ {
			// don't add the bottom right two tiles
			if x == BoardSize-1 && y >= BoardSize-2 {
				continue
			}

			val := (y+offset)%BoardSize + 1
			board.SetTileValue(x, y, val)
		}
	}

	// bad move, need to backtrack 1 and then play two final moves
	board.SetTileValue(8, 7, 8)

	board.Solve(3)

	allTilesPopulated, solvedMessage := board.allTilesPopulated()
	if !allTilesPopulated {
		t.Errorf("Board is not solved with message '%v'", solvedMessage)
	}

	isValid, message := board.isValid()
	if !isValid {
		t.Errorf("Expected board to be valid but got invalid with message '%v'", message)
	}
}
