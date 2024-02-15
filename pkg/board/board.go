package board

type Board struct {
	tiles [2][2]int
}

func New() Board {
	tiles := [2][2]int{{1, 2}, {3, 4}}

	return Board{
		tiles,
	}
}
func (board Board) String() string {
	return "board string"
}
