package app

type Board = [9][9]uint8

type Checker struct {
	board Board
}

func NewChecker(board Board) *Checker {
	return &Checker{
		board: board,
	}
}

func (c *Checker) IsValidHorizontalLine(line uint8) bool {
	return hasOneToNine(func(expectedNumber uint8) bool {
		for _, actualANumber := range c.board[line] {
			if actualANumber == expectedNumber {
				return true
			}
		}

		return false
	})
}

func (c *Checker) IsValidVerticalLine(line uint8) bool {
	return hasOneToNine(func(expectedNumber uint8) bool {
		for _, currentActualRow := range c.board {
			if currentActualRow[line] == expectedNumber {
				return true
			}
		}

		return false
	})
}

func (c *Checker) IsValidNumberBlock(position uint8) bool {
	targetPositions := getNumberBlockPositions(position)

	const ROW_INDEX = 0
	const COLUMN_INDEX = 1

	return hasOneToNine(func(expectedNumber uint8) bool {
		for _, currentPosition := range targetPositions {
			if c.board[currentPosition[ROW_INDEX]][currentPosition[COLUMN_INDEX]] == expectedNumber {
				return true
			}
		}

		return false
	})
}

func getNumberBlockPositions(target uint8) [9][2]uint8 {
	firstPositions := [9][2]uint8{
		{0, 0}, {0, 3}, {0, 6},
		{3, 0}, {3, 3}, {3, 6},
		{6, 0}, {6, 3}, {6, 6},
	}

	top := firstPositions[target][0]
	left := firstPositions[target][1]

	return [9][2]uint8{
		{top, left}, {top, left + 1}, {top, left + 2},
		{top + 1, left}, {top + 1, left + 1}, {top + 1, left + 2},
		{top + 2, left}, {top + 2, left + 1}, {top + 2, left + 2},
	}
}

func hasOneToNine(has func(uint8) bool) bool {
	expected := [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, expectedNumber := range expected {
		result := has(expectedNumber)

		if !result {
			return false
		}
	}

	return true
}
