package app

type Resolver struct {
	board      Board
	isComplete bool
}

func NewResolver(board Board) *Resolver {
	return &Resolver{
		board:      board,
		isComplete: false,
	}
}

func (r *Resolver) Resolve() Resolver {
	emptySpaces := r.board.SearchNotYetEntered()

	for _, input := range []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		filled := r.board.FillIn(emptySpaces[0][0], emptySpaces[0][1], input)

		if NewChecker(filled).IsComplete() {
			return Resolver{
				board:      r.board,
				isComplete: true,
			}
		}
	}

	return Resolver{
		board:      r.board,
		isComplete: false,
	}
}

func (r *Resolver) IsComplete() bool {
	return r.isComplete
}
