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
	emptySpaces := r.board.SearchNotYetEntered()

	for _, input := range []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		r.board = r.board.FillIn(emptySpaces[0][0], emptySpaces[0][1], input)

		if NewChecker(r.board).IsComplete() {
			r.IsComplete = true
		}
	}
}
