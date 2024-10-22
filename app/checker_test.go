package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
