package board

import (
	"reflect"
	"testing"
)

func TestPushPop(t *testing.T) {
	h := HistoryStack{}
	tile1 := NewTile(0, 0)
	tile2 := NewTile(0, 1)
	h.push(&tile1)
	h.push(&tile2)

	if len(h.tiles) != 2 {
		t.Errorf("Expected history to have len 2, got %v", len(h.tiles))
	}

	poppedTile1, empty := h.pop()

	if empty || h.isEmpty() {
		t.Errorf("Expected history to not be empty after first pop")
	}

	if reflect.DeepEqual(tile1, poppedTile1) {
		t.Errorf("Expected popped tile to be tile1 got %v", poppedTile1)
	}

	poppedTile2, empty := h.pop()

	if !empty || !h.isEmpty() {
		t.Errorf("Expected history to be empty after second pop")
	}

	if reflect.DeepEqual(tile2, poppedTile2) {
		t.Errorf("Expected popped tile to be tile2 got %v", poppedTile1)
	}
}
