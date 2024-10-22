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
