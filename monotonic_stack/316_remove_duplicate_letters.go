package monotonic_stack

// Given a string s,
// remove duplicate letters so that every letter appears once and only once.
// You must make sure your result is the smallest in lexicographical order among all possible results.
//
//Example 1:
//Input: s = "bcabc"
//Output: "abc"
//
//Example 2:
//Input: s = "cbacdcbc"
//Output: "acdb"

func removeDuplicateLetters(s string) string {
	// to count how many letters remain in the left
	counter := [26]int{}
	for _, ch := range s {
		counter[ch-'a']++
	}

	var stack []byte
	var inStack [26]bool
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			// the current character is smaller then the stack top element
			// still have the same character as the top element on the left
			for len(stack) > 0 && ch < stack[len(stack)-1] && counter[stack[len(stack)-1]-'a'] > 0 {
				inStack[stack[len(stack)-1]-'a'] = false
				stack = stack[:len(stack)-1] // pop element
			}
			inStack[ch-'a'] = true
			stack = append(stack, ch)
		}
		counter[ch-'a']--
	}

	return string(stack)
}
