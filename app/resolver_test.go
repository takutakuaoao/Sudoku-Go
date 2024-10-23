package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolve(t *testing.T) {
	sut := NewResolver(
		*NewBoard(
			[9][9]uint8{
				{0, 1, 9, 5, 4, 3, 6, 7, 8},
				{5, 4, 3, 8, 7, 6, 9, 1, 2},
				{8, 7, 6, 2, 1, 9, 3, 4, 5},
				{4, 3, 2, 7, 6, 5, 8, 9, 1},
				{7, 6, 5, 1, 9, 8, 2, 3, 4},
				{1, 9, 8, 4, 3, 2, 5, 6, 7},
				{3, 2, 1, 6, 5, 4, 7, 8, 9},
				{6, 5, 4, 9, 8, 7, 1, 2, 3},
				{9, 8, 7, 3, 2, 1, 4, 5, 6},
			},
		),
	)

	sut.resolve()

	assert.Equal(t, true, sut.IsComplete)
}
