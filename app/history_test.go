package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	sut := NewHistory([][2]uint8{{1, 0}})

	y, x := sut.GetPosition()

	assert.Equal(t, uint8(1), y)
	assert.Equal(t, uint8(0), x)
}
