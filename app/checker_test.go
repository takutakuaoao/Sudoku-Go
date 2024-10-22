package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			stu := NewChecker(tt.list)

			assert.Equal(t, tt.expected, stu.IsValidHorizontalLine(0))
		})
	}
}
