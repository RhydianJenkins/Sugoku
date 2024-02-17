package board

import (
	"fmt"
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

// TODO
func TestSolveOneStep(t *testing.T) {
	board := New()
	solveOneStep(&board)
	solveOneStep(&board)
	solveOneStep(&board)

	// is random /really/ random?
	fmt.Println(board)
}
