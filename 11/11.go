package main

import (
	"fmt"
)

func maxArea(height []int) int {
	var max int = 0
	var leftIndex int = 0
	var rightIndex int = len(height) - 1
	for leftIndex < rightIndex {
		fmt.Println(leftIndex, rightIndex)
		max = maximum(minimum(height[leftIndex], height[rightIndex])*(rightIndex-leftIndex), max)
		if height[leftIndex] < height[rightIndex] {
			leftIndex++
			continue
		}
		rightIndex--
	}
	return max
}

func maximum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimum(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	array := []int{2, 3, 1, 9, 10, 2, 5}
	fmt.Println(maxArea(array))
}
