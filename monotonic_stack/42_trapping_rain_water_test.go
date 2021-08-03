package monotonic_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrap(t *testing.T) {
	tests := []struct{
		input []int
		want int
	} {
		{
			input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:  6,
		},
		{
			input: []int{4, 2, 0, 3, 2, 5},
			want:  9,
		},
	}

	for _, tt := range tests {
		got := trap(tt.input)
		assert.Equal(t, tt.want, got)
	}
}
