package app

func NewHistory(positions [][2]uint8) *History {
	return &History{
		positions: positions,
	}
}

type History struct {
	positions [][2]uint8
}

func (h *History) GetPosition() (uint8, uint8) {
	return h.positions[0][0], h.positions[0][1]
}
