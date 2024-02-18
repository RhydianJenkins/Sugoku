package board

import (
	"reflect"
	"testing"
)

func TestGetCol(t *testing.T) {
	colNum := BoardSize - 1
	b := New()
	col := b.GetCol(colNum)
	tiles := b.GetTiles()

	for i := 0; i < BoardSize; i++ {
		if col[i].x != tiles[i][colNum].x || col[i].y != tiles[i][colNum].y {
			t.Errorf("Expected 0, got %v", col[i])
		}
	}
}

func TestCalculatePossibleValuesWithNonEmptyVal(t *testing.T) {
	returned := calculatePossibleValues(New(), 0, 0)
	expected := []int{}

	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("Expected %v, got %v", expected, returned)
	}
}

func TestCalculatePossibleValuesWithEmptyVal(t *testing.T) {
	returned := calculatePossibleValues(New(), 0, 1)
	expected := []int{2, 3, 4}

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
	board := New()
	returned := board.findLowestEntropyTiles()
	expected := []*Tile{board.GetTile(0, 1), board.GetTile(1, 0)}

	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("Expected %v, got %v", expected, returned)
	}
}

func TestSolveOneStep(t *testing.T) {
	board := New()
	numSolvedAtStart := 1 // TODO change when revisiting board initialisation
	for i := 0; i < NumValues-numSolvedAtStart; i++ {
		solveOneStep(&board)
	}

	isValid, message := boardIsValid(board)
	if !isValid {
		t.Errorf("Expected board to be valid, got invalid with message %v", message)
	}
}

func TestBoardIsValid(t *testing.T) {
	board := New()
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
