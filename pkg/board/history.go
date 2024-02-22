package board

type HistoryStack struct {
	tiles []*Tile
}

func (h *HistoryStack) push(tile *Tile) {
	h.tiles = append(h.tiles, tile)
}

func (h *HistoryStack) pop() (tile *Tile, empty bool) {
	if h.isEmpty() {
		return &Tile{}, true
	}

	numTiles := len(h.tiles)
	poppedTile := h.tiles[numTiles-1]
	h.tiles = h.tiles[:numTiles-1]

	return poppedTile, false
}

func (h HistoryStack) isEmpty() bool {
	return len(h.tiles) == 0
}
