package app

type Checker struct {
	board Board
}

func NewCheckerFromArray(board [9][9]uint8) *Checker {
	return NewChecker(*NewBoard(board))
}

func NewChecker(board Board) *Checker {
	return &Checker{
		board: board,
	}
}

func (c *Checker) IsComplete() bool {
	targets := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8}

	for _, target := range targets {
		if !c.IsValidHorizontalLine(target) {
			return false
		}

		if !c.IsValidHorizontalLine(target) {
			return false
		}

		if !c.IsValidNumberBlock(target) {
			return false
		}
	}

	return true
}

func (c *Checker) IsValidHorizontalLine(line uint8) bool {
	return hasOneToNine(func(expectedNumber uint8) bool {
		return c.board.HasInRow(line, expectedNumber)
	})
}

func (c *Checker) IsValidVerticalLine(line uint8) bool {
	return hasOneToNine(func(expectedNumber uint8) bool {
		return c.board.HasInColumn(line, expectedNumber)
	})
}

func (c *Checker) IsValidNumberBlock(position uint8) bool {
	targetPositions := getNumberBlockPositions(position)

	const ROW_INDEX = 0
	const COLUMN_INDEX = 1

	return hasOneToNine(func(expectedNumber uint8) bool {
		for _, currentPosition := range targetPositions {
			if c.board.Has(currentPosition[ROW_INDEX], currentPosition[COLUMN_INDEX], expectedNumber) {
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
