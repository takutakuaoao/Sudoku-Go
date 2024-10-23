package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveLastNumberSpace(t *testing.T) {
	cases := []struct {
		name  string
		board [9][9]uint8
	}{
		{
			name: "empty 0, 0",
			board: [9][9]uint8{
				{NOT_YET_INPUT, 1, 9, 5, 4, 3, 6, 7, 8},
				{5, 4, 3, 8, 7, 6, 9, 1, 2},
				{8, 7, 6, 2, 1, 9, 3, 4, 5},
				{4, 3, 2, 7, 6, 5, 8, 9, 1},
				{7, 6, 5, 1, 9, 8, 2, 3, 4},
				{1, 9, 8, 4, 3, 2, 5, 6, 7},
				{3, 2, 1, 6, 5, 4, 7, 8, 9},
				{6, 5, 4, 9, 8, 7, 1, 2, 3},
				{9, 8, 7, 3, 2, 1, 4, 5, 6},
			},
		},
		{
			name: "empty 0, 7",
			board: [9][9]uint8{
				{2, 1, 9, 5, 4, 3, 6, NOT_YET_INPUT, 8},
				{5, 4, 3, 8, 7, 6, 9, 1, 2},
				{8, 7, 6, 2, 1, 9, 3, 4, 5},
				{4, 3, 2, 7, 6, 5, 8, 9, 1},
				{7, 6, 5, 1, 9, 8, 2, 3, 4},
				{1, 9, 8, 4, 3, 2, 5, 6, 7},
				{3, 2, 1, 6, 5, 4, 7, 8, 9},
				{6, 5, 4, 9, 8, 7, 1, 2, 3},
				{9, 8, 7, 3, 2, 1, 4, 5, 6},
			},
		},
		{
			name: "empty 3, 4",
			board: [9][9]uint8{
				{2, 1, 9, 5, 4, 3, 6, 7, 8},
				{5, 4, 3, 8, 7, 6, 9, 1, 2},
				{8, 7, 6, 2, 1, 9, 3, 4, 5},
				{4, 3, 2, 7, NOT_YET_INPUT, 5, 8, 9, 1},
				{7, 6, 5, 1, 9, 8, 2, 3, 4},
				{1, 9, 8, 4, 3, 2, 5, 6, 7},
				{3, 2, 1, 6, 5, 4, 7, 8, 9},
				{6, 5, 4, 9, 8, 7, 1, 2, 3},
				{9, 8, 7, 3, 2, 1, 4, 5, 6},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			sut := NewResolver(*NewBoard(tt.board))

			sut.resolve()

			assert.Equal(t, true, sut.IsComplete)
		})
	}
}
