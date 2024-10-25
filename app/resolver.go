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
	history := *NewHistoryFromBoard(r.board)
	checker := NewChecker(&r.board)

	for !checker.IsComplete() {
		position, number, _ := history.GetNextInput()

		r.board = r.board.FillIn(position[0], position[1], number)

		if checker.OkAllRulesSpecifiedSquare(position) {
			history = history.OK()
		} else {
			history = history.NG()
			r.board = history.FillInAsUnentered(r.board)
		}
	}

	return Resolver{
		board:      r.board,
		isComplete: true,
	}
}

func (r *Resolver) IsComplete() bool {
	return r.isComplete
}
