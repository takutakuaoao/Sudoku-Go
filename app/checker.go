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
