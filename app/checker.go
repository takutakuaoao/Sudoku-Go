package app

type Checker struct {
	board [9][9]uint8
}

func NewChecker(board [9][9]uint8) *Checker {
	return &Checker{
		board: board,
	}
}

func (c *Checker) IsValidHorizontalLine(line uint8) bool {
	expected := [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, expectedNumber := range expected {
		isMatch := false

		for _, actualNumber := range c.board[line] {
			if actualNumber == expectedNumber {
				isMatch = true
			}
		}

		if !isMatch {
			return false
		}
	}

	return true
}
