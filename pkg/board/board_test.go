package board

import (
	"reflect"
	"testing"
)

func createBoard() Board {
	b := NewBoard([]TileVal{
		TileVal{0, 0, 1},
	})

	return b
}

func TestGetCol(t *testing.T) {
	colNum := BoardSize - 1
	b := createBoard()
	col := b.GetCol(colNum)
	tiles := b.GetTiles()

	for i := 0; i < BoardSize; i++ {
		if col[i].x != tiles[i][colNum].x || col[i].y != tiles[i][colNum].y {
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
	board := NewBoard([]TileVal{
		TileVal{0, 0, 1},
		TileVal{1, 1, 2},
		TileVal{4, 4, 2},
		TileVal{5, 4, 5},
	})
	numReturned := len(board.findLowestEntropyTiles())
	numExpected := 1

	if numReturned != numExpected {
		t.Errorf("Expected %v, but got %v", numExpected, numReturned)
	}

	lowest := board.findLowestEntropyTiles()[0]
	if lowest.x != 0 || lowest.y != 4 {
		t.Errorf("Expected lowest tile to be (0, 4) but got (%v, %v)", lowest.x, lowest.y)
	}

	expected := []int{3, 4, 6, 7, 8, 9}
	if !reflect.DeepEqual(expected, lowest.possibleValues) {
		t.Errorf("Expected possible values to be %v, but got %v", expected, lowest.possibleValues)
	}
}

func TestFindLowestEntropyTilesOnEmptyBoard(t *testing.T) {
	board := NewBoard([]TileVal{})
	numReturned := len(board.findLowestEntropyTiles())
	numExpected := BoardSize * BoardSize

	if numReturned != numExpected {
		t.Errorf("Expected %v, got %v", numExpected, numReturned)
	}
}

func TestSolveOneStep(t *testing.T) {
	board := createBoard()
	for i := 0; i < BoardSize-board.numPrePopulatedTiles; i++ {
		solveOneStep(&board)
	}

	isValid, message := boardIsValid(board)

	if !isValid {
		t.Errorf("Expected board to be valid, got invalid with message %v", message)
	}
}

func TestBoardIsValid(t *testing.T) {
	board := createBoard()
	isValid, message := boardIsValid(board)
	if !isValid {
		t.Errorf("Expected starting board to be valid, got invalid with message %v", message)
	}

	board.GetTile(0, 1).value = 1
	isValid, message = boardIsValid(board)
	if isValid {
		t.Errorf("Expected modified board to be invalid, got valid with message %v", message)
	}
}

func TestGetBlock(t *testing.T) {
	board := createBoard()
	board.GetTile(3, 4).value = 2
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
