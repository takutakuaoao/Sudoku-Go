package app

func NewHistory(positions [][2]uint8) *History {
	return &History{
		positions:       positions,
		currentPosition: positions[0],
		currentNumber:   1,
	}
}

type History struct {
	positions       [][2]uint8
	currentPosition [2]uint8
	currentNumber   uint8
}

func (h History) NG() History {
	return History{
		positions:       h.positions,
		currentPosition: h.currentPosition,
		currentNumber:   h.currentNumber + 1,
	}
}

func (h History) GetInput() ([2]uint8, uint8) {
	return h.currentPosition, h.currentNumber
}
