package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	sut := NewHistory([][2]uint8{{1, 0}})

	position, number := sut.GetInput()

	assert.Equal(t, [2]uint8{1, 0}, position)
	assert.Equal(t, uint8(1), number)
}

func TestNG(t *testing.T) {
	sut := NewHistory([][2]uint8{{2, 0}})

	position, number := sut.NG().GetInput()

	assert.Equal(t, position, [2]uint8{2, 0})
	assert.Equal(t, number, uint8(2))
}
