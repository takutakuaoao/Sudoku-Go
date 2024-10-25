package app

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
	if h.getCurrent().hasLast() {
		return h.backPosition()
	}

	return h.nextNumber()
}

func (h History) GetInput() ([2]uint8, uint8, bool) {
	if h.currentIndex+1 > uint8(len(h.positions)) {
		return [2]uint8{}, 0, true
	}

	position, lastNumber := h.getCurrent().toPrimitive()

	return position, lastNumber, false
}

func (h History) nextNumber() History {
	return h.updatePosition(h.currentIndex, h.getCurrent().nextNumber())
}

func (h History) nextPosition() History {
	return h.updateIndex(h.currentIndex + 1)
}

func (h History) backPosition() History {
	tmp := h.updatePosition(h.currentIndex, h.getCurrent().resetNumber())

	if tmp.positions[h.currentIndex-1].hasLast() {
		tmp = tmp.updatePosition(h.currentIndex-1, h.positions[h.currentIndex-1].resetNumber())
	} else {
		tmp = tmp.updatePosition(h.currentIndex-1, h.positions[h.currentIndex-1].nextNumber())
	}

	return tmp.updateIndex(h.currentIndex - 1)
}

func (h History) getCurrent() position {
	return h.positions[h.currentIndex]
}

func (h History) updatePosition(index uint8, newPos position) History {
	tmpPositions := h.positions
	tmpPositions[index] = newPos

	return History{
		currentIndex: h.currentIndex,
		positions:    tmpPositions,
	}
}

func (h History) updateIndex(index uint8) History {
	return History{
		positions:    h.positions,
		currentIndex: index,
	}
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
