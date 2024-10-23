package app

func NewHistory(positions [][2]uint8) *History {
	_positions := []position{}

	for _, _position := range positions {
		_positions = append(_positions, *newPosition(_position))
	}

	return &History{
		currentIndex: 0,
		positions:    _positions,
	}
}

type History struct {
	currentIndex uint8
	positions    []position
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

func (h History) GetInput() ([2]uint8, uint8) {
	return h.getCurrent().toPrimitive()
}

func (h History) nextNumber() History {
	return h.updatePosition(h.currentIndex, h.getCurrent().nextNumber())
}

func (h History) nextPosition() History {
	return h.updateIndex(h.currentIndex + 1)
}

func (h History) backPosition() History {
	return h.
		updatePosition(h.currentIndex-1, h.getPrevPosition().nextNumber()).
		updateIndex(h.currentIndex - 1)
}

func (h History) getCurrent() position {
	return h.positions[h.currentIndex]
}

func (h History) getPrevPosition() position {
	return h.positions[h.currentIndex-1]
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
	lastNumber uint8
}

func newPosition(pos [2]uint8) *position {
	return &position{
		y:          pos[0],
		x:          pos[1],
		lastNumber: 1,
	}
}

func (p position) nextNumber() position {
	return position{
		y:          p.y,
		x:          p.x,
		lastNumber: p.lastNumber + 1,
	}
}

func (p position) hasLast() bool {
	return p.lastNumber == 9
}

func (p position) toPrimitive() ([2]uint8, uint8) {
	return [2]uint8{p.y, p.x}, p.lastNumber
}
