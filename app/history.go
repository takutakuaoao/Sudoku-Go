package main

func NewHistoryFromBoard(board Board) *History {
	noEnteredSquares := board.SearchNotYetEntered()
	_positions := []position{}

	for _, pos := range noEnteredSquares {
		candidates := board.NarrowDownEnterableNumbers(pos)
		_positions = append(_positions, *newPosition(pos, candidates))
	}

	return &History{
		currentIndex: 0,
		positions:    _positions,
	}
}

func NewHistory(positions [][2]uint8) *History {
	_positions := []position{}

	for _, _position := range positions {
		_positions = append(_positions, *newPosition(_position, []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	}

	return &History{
		currentIndex: 0,
		positions:    _positions,
	}
}

type History struct {
	currentIndex   uint8
	positions      []position
	resetPositions [][2]uint8
}

func (h History) OK() History {
	return h.nextPosition()
}

func (h History) NG() History {
	tmp := h.cleanResetPositions().
		addResetPosition(h.getCurrent())

	if tmp.getCurrent().hasLast() {
		return tmp.backPosition()
	}

	return tmp.nextNumber()
}

func (h History) GetNextInput() ([2]uint8, uint8, bool) {
	if h.currentIndex+1 > uint8(len(h.positions)) {
		return [2]uint8{}, 0, true
	}

	position, lastNumber := h.getCurrent().toPrimitive()

	return position, lastNumber, false
}

func (h History) FillInAsUnentered(board Board) Board {
	tmp := board

	for _, resetPosition := range h.resetPositions {
		tmp = tmp.FillIn(resetPosition[0], resetPosition[1], 0)
	}

	return tmp
}

func (h History) cleanResetPositions() History {
	return History{
		currentIndex:   h.currentIndex,
		positions:      h.positions,
		resetPositions: [][2]uint8{},
	}
}

func (h History) nextNumber() History {
	return h.updatePosition(h.currentIndex, h.getCurrent().nextNumber())
}

func (h History) nextPosition() History {
	return h.updateIndex(h.currentIndex + 1)
}

func (h History) getCurrent() position {
	return h.positions[h.currentIndex]
}

func (h History) updatePosition(index uint8, newPos position) History {
	tmpPositions := h.positions
	tmpPositions[index] = newPos

	return History{
		currentIndex:   h.currentIndex,
		positions:      tmpPositions,
		resetPositions: h.resetPositions,
	}
}

func (h History) updateIndex(index uint8) History {
	return History{
		positions:      h.positions,
		currentIndex:   index,
		resetPositions: h.resetPositions,
	}
}

func (h History) addResetPosition(position position) History {
	pos, _ := position.toPrimitive()

	tmp := h.resetPositions
	tmp = append(tmp, pos)

	return History{
		currentIndex:   h.currentIndex,
		positions:      h.positions,
		resetPositions: tmp,
	}
}

func (h History) backPosition() History {
	tmp := h.updatePosition(h.currentIndex, h.getCurrent().resetNumber())

	if tmp.getPrevPosition().hasLast() {
		return tmp.
			backPrevFirstNumber().
			backPosition()
	}

	return tmp.nextPrevPositionNumber()
}

func (h History) backPrevFirstNumber() History {
	return h.
		updatePrevPosition(h.getPrevPosition().resetNumber()).
		addResetPosition(h.getPrevPosition())
}

func (h History) nextPrevPositionNumber() History {
	return h.updatePrevPosition(h.getPrevPosition().nextNumber())
}

func (h History) updatePrevPosition(prevPosition position) History {
	return h.updatePosition(h.currentIndex-1, prevPosition).
		updateIndex(h.currentIndex - 1)
}

func (h History) getPrevPosition() position {
	return h.positions[h.currentIndex-1]
}

type position struct {
	y          uint8
	x          uint8
	lastIndex  uint8
	candidates []uint8
}

func newPosition(pos [2]uint8, candidates []uint8) *position {
	return &position{
		y:          pos[0],
		x:          pos[1],
		lastIndex:  0,
		candidates: candidates,
	}
}

func (p position) nextNumber() position {
	return position{
		y:          p.y,
		x:          p.x,
		lastIndex:  p.lastIndex + 1,
		candidates: p.candidates,
	}
}

func (p position) hasLast() bool {
	return int(p.lastIndex) == len(p.candidates)-1
}

func (p position) toPrimitive() ([2]uint8, uint8) {
	return [2]uint8{p.y, p.x}, p.candidates[p.lastIndex]
}

func (p position) resetNumber() position {
	return position{
		y:          p.y,
		x:          p.x,
		lastIndex:  0,
		candidates: p.candidates,
	}
}
