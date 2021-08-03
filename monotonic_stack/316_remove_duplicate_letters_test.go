package monotonic_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "cbacdcbc",
			want:  "acdb",
		},
		{
			input: "bcabc",
			want:  "abc",
		},
		{
			input: "bcbcbc",
			want: "bc",
		},
		{
			input: "dbcdabb",
			want:  "bcda",
		},
	}

	for _, tt := range tests {
		got := removeDuplicateLetters(tt.input)
		assert.Equal(t, tt.want, got)
	}
}
