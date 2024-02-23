package board

type HistoryStack struct {
	tiles []*Tile
}

func (h *HistoryStack) push(tile *Tile) {
	h.tiles = append(h.tiles, tile)
}

func (h *HistoryStack) pop() *Tile {
	if h.isEmpty() {
		return nil
	}

	numTiles := len(h.tiles)
	poppedTile := h.tiles[numTiles-1]
	h.tiles = h.tiles[:numTiles-1]

	return poppedTile
}

func (h HistoryStack) isEmpty() bool {
	return len(h.tiles) == 0
}
