package main

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

func (b *Board) HasInBlock(position [2]uint8, value uint8) bool {
	blockPositions := NewBlock().GetAllPositionInBlock(position)

	for _, blockPos := range blockPositions {
		if b.Has(blockPos[0], blockPos[1], value) {
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

func (b *Board) GetNumbersInBlock(position [2]uint8) [9]uint8 {
	result := []uint8{}

	for _, pos := range NewBlock().GetAllPositionInBlock(position) {
		result = append(result, b.GetPositionNumber(pos))
	}

	return [9]uint8(result)
}

func (b *Board) NarrowDownEnterableNumbers(position [2]uint8) []uint8 {
	numbers := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := []uint8{}

	for _, number := range numbers {
		if b.HasInRow(position[0], number) || b.HasInColumn(position[1], number) || b.HasInBlock(position, number) {
			continue
		}

		result = append(result, number)
	}

	return result
}
