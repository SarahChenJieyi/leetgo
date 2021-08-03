package monotonic_stack

import "math"

func lengthOfLongestSubstring(s string) int {
	subMap := make(map[string]int, 0)
	var result float64
	for x, vStart := range s {
		subMap[string(vStart)] = 1
		for _, vEnd := range s[x+1:] {
			if _, ok := subMap[string(vEnd)]; ok {
				break
			}
			subMap[string(vEnd)] = 1
		}
		result = math.Max(result, float64(len(subMap)))
		subMap = make(map[string]int, 0)
	}
	return int(result)
}
