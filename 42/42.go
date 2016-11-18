package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	len_height := len(height)
	if len_height <= 2 {
		return 0
	}

	var left []int = make([]int, len_height)
	var right []int = make([]int, len_height)

	left[0] = height[0]
	right[len_height-1] = height[len_height-1]
	for i := 1; i < len_height; i++ {
		left[i] = max(left[i-1], height[i])
		right[len_height-1-i] = max(right[len_height-i], height[len_height-1-i])
	}

	var water int = 0
	for i := 0; i < len_height; i++ {
		water += min(left[i], right[i]) - height[i]
	}

	return water
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
