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
		for columnIndex := range row {
			if b.IsUnEntered([2]uint8{uint8(rowIndex), uint8(columnIndex)}) {
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

func (b *Board) DuplicateNumberInRow(rowNumber uint8, number uint8) bool {
	count := 0

	for _, value := range b.rows[rowNumber] {
		if value == number {
			count++
		}
	}

	return count >= 2
}

func (b *Board) DuplicateNumberInColumn(columnNumber uint8, number uint8) bool {
	count := 0

	for _, row := range b.rows {
		if row[columnNumber] == number {
			count++
		}
	}

	return count >= 2
}

func (b *Board) GetPositionNumber(position [2]uint8) uint8 {
	return b.rows[position[0]][position[1]]
}

func (b *Board) IsUnEntered(position [2]uint8) bool {
	return b.GetPositionNumber(position) == NOT_YET_INPUT
}
