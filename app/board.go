package app

const NOT_YET_INPUT = 0

type Board struct {
	rows [9][9]uint8
}

func NewBoard(rows [9][9]uint8) *Board {
	return &Board{
		rows: rows,
	}
}

func (b *Board) FillIn(row uint8, column uint8, value uint8) Board {
	b.rows[row][column] = value

	return *NewBoard(b.rows)
}

func (b *Board) SearchNotYetEntered() [][2]uint8 {
	result := [][2]uint8{}

	for rowIndex, row := range b.rows {
		for columnIndex, column := range row {
			if column == NOT_YET_INPUT {
				result = append(result, [2]uint8{uint8(rowIndex), uint8(columnIndex)})
			}
		}
	}

	return result
}

func (b *Board) HasInRow(row uint8, value uint8) bool {
	for columnIndex := range b.rows[row] {
		if b.Has(row, uint8(columnIndex), value) {
			return true
		}
	}

	return false
}

func (b *Board) HasInColumn(column uint8, value uint8) bool {
	for rowIndex := range b.rows {
		if b.Has(uint8(rowIndex), column, value) {
			return true
		}
	}

	return false
}

func (b *Board) Has(row uint8, column uint8, value uint8) bool {
	return b.rows[row][column] == value
}
