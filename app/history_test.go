package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	sut := NewHistory([][2]uint8{{1, 0}})

	position, number, _ := sut.GetInput()

	assert.Equal(t, [2]uint8{1, 0}, position)
	assert.Equal(t, uint8(1), number)
}

func Test_return_next_number_if_number_other_then_the_last_is_NG(t *testing.T) {
	sut := NewHistory([][2]uint8{{2, 0}})

	position, number, _ := sut.NG().GetInput()

	assert.Equal(t, position, [2]uint8{2, 0})
	assert.Equal(t, number, uint8(2))
}

func Test_if_the_last_number_nine_is_NG_the_history_return_prev_position_and_next_number_of_prev_success_number(t *testing.T) {
	sut := NewHistory([][2]uint8{{2, 0}, {3, 1}})

	position, number, _ := sut.OK().NG().NG().NG().NG().NG().NG().NG().NG().NG().GetInput()

	assert.Equal(t, position, [2]uint8{2, 0})
	assert.Equal(t, number, uint8(2))
}

func Test_if_all_squares_are_OK(t *testing.T) {
	sut := NewHistory([][2]uint8{{1, 0}, {3, 3}, {4, 4}})

	_, _, isAllOK := sut.OK().OK().OK().GetInput()

	assert.True(t, isAllOK)
}

func Test_to_resume_from_a_successful_square_if_the_squares_are_not_filled_consecutively(t *testing.T) {
	sut := History{
		currentIndex: 2,
		positions: []position{
			{
				y:          0,
				x:          1,
				lastNumber: 1,
			},
			{
				y:          0,
				x:          2,
				lastNumber: 9,
			},
			{
				y:          0,
				x:          3,
				lastNumber: 9,
			},
		},
	}

	position, number, _ := sut.NG().GetInput()

	assert.Equal(t, [2]uint8{0, 1}, position)
	assert.Equal(t, uint8(2), number)
}