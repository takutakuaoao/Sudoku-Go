package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOkIfBoardIsComplete(t *testing.T) {
	sut := NewChecker(Board{
		{2, 1, 9, 5, 4, 3, 6, 7, 8},
		{5, 4, 3, 8, 7, 6, 9, 1, 2},
		{8, 7, 6, 2, 1, 9, 3, 4, 5},
		{4, 3, 2, 7, 6, 5, 8, 9, 1},
		{7, 6, 5, 1, 9, 8, 2, 3, 4},
		{1, 9, 8, 4, 3, 2, 5, 6, 7},
		{3, 2, 1, 6, 5, 4, 7, 8, 9},
		{6, 5, 4, 9, 8, 7, 1, 2, 3},
		{9, 8, 7, 3, 2, 1, 4, 5, 6},
	})

	assert.Equal(t, true, sut.IsComplete())
}

func TestNGIfBoardIsNotComplete(t *testing.T) {
	sut := NewChecker(Board{
		{1, 1, 9, 5, 4, 3, 6, 7, 8},
		{1, 4, 3, 8, 7, 6, 9, 1, 2},
		{1, 7, 6, 2, 1, 9, 3, 4, 5},
		{1, 3, 2, 7, 6, 5, 8, 9, 1},
		{1, 6, 5, 1, 9, 8, 2, 3, 4},
		{1, 9, 8, 4, 3, 2, 5, 6, 7},
		{1, 2, 1, 6, 5, 4, 7, 8, 9},
		{1, 5, 4, 9, 8, 7, 1, 2, 3},
		{1, 8, 7, 3, 2, 1, 4, 5, 6},
	})

	assert.Equal(t, false, sut.IsComplete())
}

func TestHorizontalANumberLine(t *testing.T) {
	cases := []struct {
		name     string
		list     Board
		expected bool
	}{
		{
			name:     "input numbers with sorted one to nine",
			list:     Board{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			expected: true,
		}, {
			name:     "input numbers with random sorted one to nine",
			list:     Board{{2, 5, 3, 6, 1, 4, 9, 8, 7}},
			expected: true,
		}, {
			name:     "input invalid numbers",
			list:     Board{{1}},
			expected: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			stu := NewChecker(tt.list)

			assert.Equal(t, tt.expected, stu.IsValidHorizontalLine(0))
		})
	}
}

func TestVerticalANumberLine(t *testing.T) {
	cases := []struct {
		name     string
		board    Board
		expected bool
	}{
		{
			name:     "success",
			board:    Board{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}},
			expected: true,
		},
		{
			name:     "invalid",
			board:    Board{{1}},
			expected: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			stu := NewChecker(tt.board)

			assert.Equal(t, tt.expected, stu.IsValidVerticalLine(0))
		})

	}
}

func TestCheckNumberBlock(t *testing.T) {
	cases := []struct {
		name        string
		board       Board
		targetBlock uint8
	}{
		{
			name: "top left block",
			board: Board{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			targetBlock: 0,
		},
		{
			name: "top middle block",
			board: Board{
				{0, 0, 0, 1, 2, 3},
				{0, 0, 0, 4, 5, 6},
				{0, 0, 0, 7, 8, 9},
			},
			targetBlock: 1,
		},
		{
			name: "top right block",
			board: Board{
				{0, 0, 0, 0, 0, 0, 1, 2, 3},
				{0, 0, 0, 0, 0, 0, 4, 5, 6},
				{0, 0, 0, 0, 0, 0, 7, 8, 9},
			},
			targetBlock: 2,
		},
		{
			name: "middle left block",
			board: Board{
				{},
				{},
				{},
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			targetBlock: 3,
		},
		{
			name: "middle middle block",
			board: Board{
				{},
				{},
				{},
				{0, 0, 0, 1, 2, 3},
				{0, 0, 0, 4, 5, 6},
				{0, 0, 0, 7, 8, 9},
			},
			targetBlock: 4,
		},
		{
			name: "middle right block",
			board: Board{
				{},
				{},
				{},
				{0, 0, 0, 0, 0, 0, 1, 2, 3},
				{0, 0, 0, 0, 0, 0, 4, 5, 6},
				{0, 0, 0, 0, 0, 0, 7, 8, 9},
			},
			targetBlock: 5,
		},
		{
			name: "bottom left block",
			board: Board{
				{},
				{},
				{},
				{},
				{},
				{},
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			targetBlock: 6,
		},
		{
			name: "bottom middle block",
			board: Board{
				{},
				{},
				{},
				{},
				{},
				{},
				{0, 0, 0, 1, 2, 3},
				{0, 0, 0, 4, 5, 6},
				{0, 0, 0, 7, 8, 9},
			},
			targetBlock: 7,
		},
		{
			name: "bottom right block",
			board: Board{
				{},
				{},
				{},
				{},
				{},
				{},
				{0, 0, 0, 0, 0, 0, 1, 2, 3},
				{0, 0, 0, 0, 0, 0, 4, 5, 6},
				{0, 0, 0, 0, 0, 0, 7, 8, 9},
			},
			targetBlock: 8,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			stu := NewChecker(tt.board)

			assert.Equal(t, true, stu.IsValidNumberBlock(tt.targetBlock))
		})
	}
}
