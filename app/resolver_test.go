package main

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

			result := sut.Resolve()

			assert.Equal(t, true, result.IsComplete())
		})
	}
}

func TestResolveMultipleNumberSpaces(t *testing.T) {
	sut := NewResolver(*NewBoard([9][9]uint8{
		{NOT_YET_INPUT, 1, 9, 5, 4, 3, 6, 7, 8},
		{5, 4, 3, 8, 7, 6, 9, 1, 2},
		{8, 7, 6, 2, 1, 9, 3, 4, 5},
		{4, 3, 2, 7, NOT_YET_INPUT, 5, 8, 9, 1},
		{7, 6, 5, 1, 9, 8, 2, 3, 4},
		{1, 9, 8, 4, 3, 2, 5, 6, 7},
		{3, 2, 1, 6, 5, 4, 7, 8, 9},
		{6, 5, 4, 9, 8, 7, 1, 2, 3},
		{9, 8, 7, 3, 2, 1, NOT_YET_INPUT, NOT_YET_INPUT, NOT_YET_INPUT},
	}))

	result := sut.Resolve()

	assert.True(t, result.IsComplete())
}

func Test_resolve_beginner(t *testing.T) {
	sut := NewResolver(*NewBoard([9][9]uint8{
		{0, 0, 1, 0, 5, 0, 0, 9, 2},
		{3, 9, 7, 2, 4, 1, 0, 6, 0},
		{0, 0, 0, 0, 3, 0, 0, 4, 0},
		{0, 0, 0, 3, 7, 2, 4, 0, 6},
		{7, 3, 0, 4, 0, 6, 2, 0, 0},
		{6, 4, 0, 0, 8, 0, 0, 0, 3},
		{0, 6, 4, 5, 0, 0, 7, 0, 9},
		{5, 0, 0, 9, 6, 4, 8, 0, 1},
		{8, 0, 9, 0, 0, 3, 0, 5, 0},
	}))

	result := sut.Resolve()

	assert.True(t, result.IsComplete())
}

func Test_resolve_easy(t *testing.T) {
	sut := NewResolver(*NewBoard([9][9]uint8{
		{3, 2, 0, 0, 0, 1, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 5},
		{0, 4, 8, 6, 9, 2, 0, 0, 1},
		{0, 7, 0, 2, 0, 0, 8, 0, 0},
		{1, 8, 2, 9, 0, 0, 0, 3, 4},
		{0, 0, 9, 0, 0, 6, 0, 1, 0},
		{0, 0, 5, 0, 0, 0, 1, 8, 6},
		{0, 0, 4, 0, 0, 0, 0, 0, 0},
		{9, 0, 0, 8, 6, 0, 4, 0, 0},
	}))

	result := sut.Resolve()

	assert.True(t, result.IsComplete())
}

func Test_have_them_solve_advanced_problem(t *testing.T) {
	sut := NewResolver(*NewBoard([9][9]uint8{
		{0, 0, 6, 0, 0, 0, 2, 0, 0},
		{0, 2, 0, 0, 0, 0, 9, 0, 8},
		{0, 8, 0, 0, 0, 0, 4, 3, 0},
		{0, 0, 5, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 9, 0, 4, 0},
		{1, 0, 0, 0, 4, 0, 6, 5, 0},
		{0, 0, 0, 9, 0, 0, 0, 0, 6},
		{0, 4, 1, 0, 3, 0, 0, 0, 0},
		{8, 0, 0, 6, 0, 0, 0, 0, 0},
	}))

	result := sut.Resolve()

	assert.True(t, result.IsComplete())
}

func Test_have_them_solve_very_hard_problem(t *testing.T) {
	sut := NewResolver(*NewBoard([9][9]uint8{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}))

	result := sut.Resolve()

	assert.True(t, result.IsComplete())
}
