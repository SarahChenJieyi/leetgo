package pointers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestLengthOfLongestSubstring(t *testing.T) {
	input := []string{"anviaj","abcabcbb","pwwkew"}
	output := []int{5,3,3}
	for i, in := range input{
		res := lengthOfLongestSubstring(in)
		fmt.Sprintln(res, output[i])
		assert.Equal(t, output[i], res)
	}
}
