package monotonic_stack

// Given n non-negative integers representing an elevation map
// where the width of each bar is 1,
// compute how much water it can trap after raining.
//
// Example 1:
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section)
// is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
// In this case, 6 units of rain water (blue section) are being trapped

func trap(height []int) int {
	var res = 0
	// will need at least 3 integers to "trap" the water
	if len(height) < 3 {
		return res
	}

	// create a monotonous stack to record the index
	var mtStack []int
	// traverse the height, put the index to the stack, when the next height is higher than the top one,
	// stop and check if can trap the water
	for i, h := range height {
		for len(mtStack) != 0 && h > height[mtStack[len(mtStack)-1]] {
			top := mtStack[len(mtStack)-1]
			mtStack = mtStack[:len(mtStack)-1]
			// length less than 2, can not trap water
			if len(mtStack) == 0 {
				break
			}
			// calculate water amount
			// trapped area = curWeight * curHeight
			// curWeight = current index - left index - 1
			// curHeight = min(current height, left height) - top height
			res += (i - mtStack[len(mtStack)-1] - 1) * (min(h, height[mtStack[len(mtStack)-1]]) - height[top])
		}
		mtStack = append(mtStack, i)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
