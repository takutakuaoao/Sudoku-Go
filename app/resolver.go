package app

type Resolver struct {
	board      Board
	IsComplete bool
}

func NewResolver(board Board) *Resolver {
	return &Resolver{
		board:      board,
		IsComplete: false,
	}
}

func (r *Resolver) resolve() {
	for _, input := range []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		r.board = r.board.FillIn(0, 0, input)

		if NewChecker(r.board).IsComplete() {
			r.IsComplete = true
		}
	}
}
