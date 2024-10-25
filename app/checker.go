package main

import (
	"errors"
)

const MAX_NUMBER_BLOCK_TARGET = 8

type Checker struct {
	board *Board
}

func NewCheckerFromArray(board [9][9]uint8) *Checker {
	return NewChecker(NewBoard(board))
}

func NewChecker(board *Board) *Checker {
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

		if result, _ := c.IsValidNumberBlock(target); !result {
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

func (c *Checker) IsValidNumberBlock(position uint8) (bool, error) {
	if position > MAX_NUMBER_BLOCK_TARGET {
		return false, errors.New(ERROR_OUT_BLOCK_NUMBER_RANGE)
	}

	targetPositions := NewBlock().GetAllPositionFromBlockNumber(position)

	const ROW_INDEX = 0
	const COLUMN_INDEX = 1

	return hasOneToNine(func(expectedNumber uint8) bool {
		for _, currentPosition := range targetPositions {
			if c.board.Has(currentPosition[ROW_INDEX], currentPosition[COLUMN_INDEX], expectedNumber) {
				return true
			}
		}

		return false
	}), nil
}

func (c *Checker) OkAllRulesSpecifiedSquare(position [2]uint8) bool {
	return c.OkVerticalSpecifiedPosition(position) &&
		c.OkHorizontalSpecifiedPosition(position) &&
		c.OkBlockSpecifiedSquare(position)
}

func (c *Checker) OkHorizontalSpecifiedPosition(position [2]uint8) bool {
	return c.tryIfNotEmptyPosition(position, func(position [2]uint8) bool {
		return !c.board.DuplicateNumberInRow(position[0], c.board.GetPositionNumber(position))
	})
}

func (c *Checker) OkVerticalSpecifiedPosition(position [2]uint8) bool {
	return c.tryIfNotEmptyPosition(position, func(position [2]uint8) bool {
		return !c.board.DuplicateNumberInColumn(
			position[1],
			c.board.GetPositionNumber(position),
		)
	})
}

func (c *Checker) OkBlockSpecifiedSquare(position [2]uint8) bool {
	return c.tryIfNotEmptyPosition(position, func(position [2]uint8) bool {
		targetNumber := c.board.GetPositionNumber(position)

		count := 0

		for _, number := range c.board.GetNumbersInBlock(position) {
			if targetNumber == number {
				count++
			}
		}

		return count == 1
	})
}

func (c *Checker) tryIfNotEmptyPosition(position [2]uint8, callBack func(position [2]uint8) bool) bool {
	if c.board.IsUnEntered(position) {
		return true
	}

	return callBack(position)
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

const ERROR_OUT_BLOCK_NUMBER_RANGE = "the number block must be between one to nine"
