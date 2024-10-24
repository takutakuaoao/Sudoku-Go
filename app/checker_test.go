package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOkIfBoardIsComplete(t *testing.T) {
	sut := NewCheckerFromArray([9][9]uint8{
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
	sut := NewCheckerFromArray([9][9]uint8{
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
		list     [9][9]uint8
		expected bool
	}{
		{
			name:     "input numbers with sorted one to nine",
			list:     [9][9]uint8{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			expected: true,
		}, {
			name:     "input numbers with random sorted one to nine",
			list:     [9][9]uint8{{2, 5, 3, 6, 1, 4, 9, 8, 7}},
			expected: true,
		}, {
			name:     "input invalid numbers",
			list:     [9][9]uint8{{1}},
			expected: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			stu := NewCheckerFromArray(tt.list)

			assert.Equal(t, tt.expected, stu.IsValidHorizontalLine(0))
		})
	}
}

func TestVerticalANumberLine(t *testing.T) {
	cases := []struct {
		name     string
		board    [9][9]uint8
		expected bool
	}{
		{
			name:     "success",
			board:    [9][9]uint8{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}},
			expected: true,
		},
		{
			name:     "invalid",
			board:    [9][9]uint8{{1}},
			expected: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			stu := NewCheckerFromArray(tt.board)

			assert.Equal(t, tt.expected, stu.IsValidVerticalLine(0))
		})

	}
}

func TestCheckNumberBlock(t *testing.T) {
	cases := []struct {
		name        string
		board       [9][9]uint8
		targetBlock uint8
	}{
		{
			name: "top left block",
			board: [9][9]uint8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			targetBlock: 0,
		},
		{
			name: "top middle block",
			board: [9][9]uint8{
				{0, 0, 0, 1, 2, 3},
				{0, 0, 0, 4, 5, 6},
				{0, 0, 0, 7, 8, 9},
			},
			targetBlock: 1,
		},
		{
			name: "top right block",
			board: [9][9]uint8{
				{0, 0, 0, 0, 0, 0, 1, 2, 3},
				{0, 0, 0, 0, 0, 0, 4, 5, 6},
				{0, 0, 0, 0, 0, 0, 7, 8, 9},
			},
			targetBlock: 2,
		},
		{
			name: "middle left block",
			board: [9][9]uint8{
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
			board: [9][9]uint8{
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
			board: [9][9]uint8{
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
			board: [9][9]uint8{
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
			board: [9][9]uint8{
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
			board: [9][9]uint8{
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
			stu := NewCheckerFromArray(tt.board)
			result, _ := stu.IsValidNumberBlock(tt.targetBlock)

			assert.Equal(t, true, result)
		})
	}
}

func TestErrorPassedOutOfNumberBlock(t *testing.T) {
	stu := NewCheckerFromArray([9][9]uint8{})

	_, err := stu.IsValidNumberBlock(9)

	assert.EqualError(t, err, ERROR_OUT_BLOCK_NUMBER_RANGE)
}

func Test_to_check_for_rule_violations_on_the_horizontal_line_of_given_square(t *testing.T) {
	cases := []struct {
		name     string
		boardRow [9]uint8
		expected bool
	}{
		{
			name:     "Duplicate numbers on a horizontal line are not acceptable.",
			boardRow: [9]uint8{1, 1, 2, 3, 4, 5, 6, 7, 8},
			expected: false,
		},
		{
			name:     "OK if there are no duplicate numbers on the horizontal line.",
			boardRow: [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: true,
		},
		{
			name:     "OK if there are no duplicate numbers on the horizontal line with some spaces.",
			boardRow: [9]uint8{1, 2, 3, 4, 0, 0, 0, 0, 0},
			expected: true,
		},
		{
			name:     "OK if the specified square is not filled in.",
			boardRow: [9]uint8{0, 1, 2, 3, 0, 0, 0, 0, 0},
			expected: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			stu := NewCheckerFromArray([9][9]uint8{
				{},
				tt.boardRow,
			})

			result := stu.OkHorizontalSpecifiedPosition([2]uint8{1, 0})

			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_to_check_for_rule_violations_on_the_vertical_line_of_given_square(t *testing.T) {
	cases := []struct {
		name         string
		verticalLine [9][9]uint8
		expected     bool
	}{
		{
			name: "Duplicate numbers on a vertical line are not acceptable.",
			verticalLine: [9][9]uint8{
				{1},
				{1},
				{2},
			},
			expected: false,
		},
		{
			name: "OK if there are no duplicate numbers on a vertical line.",
			verticalLine: [9][9]uint8{
				{1},
				{2},
			},
			expected: true,
		},
		{
			name: "OK if the specified square is not filled in.",
			verticalLine: [9][9]uint8{
				{0},
				{0},
				{0},
			},
			expected: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			stu := NewCheckerFromArray(tt.verticalLine)

			result := stu.OkVerticalSpecifiedPosition([2]uint8{0, 0})

			assert.Equal(t, tt.expected, result)
		})
	}

}
