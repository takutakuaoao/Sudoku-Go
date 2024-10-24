package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchNotYetEntered(t *testing.T) {
	sut := NewBoard([9][9]uint8{
		{NOT_YET_INPUT, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, NOT_YET_INPUT, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, NOT_YET_INPUT, NOT_YET_INPUT, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
	})

	result := sut.SearchNotYetEntered()

	assert.Equal(t, [][2]uint8{{0, 0}, {1, 1}, {2, 2}, {2, 3}}, result)
}

func Test_to_return_the_positions_of_the_block_to_which_the_specified_square_belongs(t *testing.T) {
	cases := []struct {
		name     string
		position [2]uint8
		board    [9][9]uint8
	}{
		{
			name:     "in block 1",
			position: [2]uint8{0, 1},
			board: [9][9]uint8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			name:     "in block 2",
			position: [2]uint8{0, 4},
			board: [9][9]uint8{
				{0, 0, 0, 1, 2, 3},
				{0, 0, 0, 4, 5, 6},
				{0, 0, 0, 7, 8, 9},
			},
		},
		{
			name:     "in block 3",
			position: [2]uint8{1, 7},
			board: [9][9]uint8{
				{0, 0, 0, 0, 0, 0, 1, 2, 3},
				{0, 0, 0, 0, 0, 0, 4, 5, 6},
				{0, 0, 0, 0, 0, 0, 7, 8, 9},
			},
		},
		{
			name:     "in block 4",
			position: [2]uint8{5, 0},
			board: [9][9]uint8{
				{},
				{},
				{},
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			name:     "in block 5",
			position: [2]uint8{4, 5},
			board: [9][9]uint8{
				{},
				{},
				{},
				{0, 0, 0, 1, 2, 3},
				{0, 0, 0, 4, 5, 6},
				{0, 0, 0, 7, 8, 9},
			},
		},
		{
			name:     "in block 6",
			position: [2]uint8{5, 6},
			board: [9][9]uint8{
				{},
				{},
				{},
				{0, 0, 0, 0, 0, 0, 1, 2, 3},
				{0, 0, 0, 0, 0, 0, 4, 5, 6},
				{0, 0, 0, 0, 0, 0, 7, 8, 9},
			},
		},
		{
			name:     "in block 7",
			position: [2]uint8{6, 2},
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
		},
		{
			name:     "in block 8",
			position: [2]uint8{7, 3},
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
		},
		{
			name:     "in block 9",
			position: [2]uint8{8, 7},
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
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			sut := NewBoard(tt.board)

			result := sut.GetNumbersInBlock(tt.position)

			assert.Equal(t, [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
		})
	}
}

func Test_to_narrow_down_the_numbers_that_can_be_entered(t *testing.T) {
	sut := NewBoard([9][9]uint8{
		{0, 0, 1, 0, 5, 0, 0, 9, 2},
		{3, 9, 7, 2, 4, 1, 0, 6, 0},
		{0, 0, 0, 0, 3, 0, 0, 4, 0},
		{0, 0, 0, 3, 7, 2, 4, 0, 6},
		{7, 3, 0, 4, 0, 6, 2, 0, 0},
		{6, 4, 0, 0, 8, 0, 0, 0, 3},
		{0, 6, 4, 5, 0, 0, 7, 0, 9},
		{5, 0, 0, 9, 6, 4, 8, 0, 1},
		{8, 0, 9, 0, 0, 3, 0, 5, 0},
	})

	result := sut.NarrowDownEnterableNumbers([2]uint8{0, 3})
	assert.Equal(t, []uint8{6, 7, 8}, result)
}
