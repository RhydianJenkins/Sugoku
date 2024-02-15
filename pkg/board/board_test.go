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
		if col[i] != tiles[i][colNum] {
			t.Errorf("Expected 0, got %v", col[i])
		}
	}
}

func TestCalculatePossibleValuesWithNonEmptyVal(t *testing.T) {
	b := New()
	returned := b.CalculatePossibleValues(0, 0)
	expected := []int{1}

	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("Expected %v, got %v", expected, returned)
	}
}

func TestCalculatePossibleValuesWithEmptyVal(t *testing.T) {
	b := New()
	returned := b.CalculatePossibleValues(0, 1)
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
