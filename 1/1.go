package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	fmt.Println(len(nums))
	var m map[int]int = make(map[int]int, 1024*1024)
	for i, n := range nums {
		m[n] = i
	}

	for i, n := range nums {
		if index, ok := m[target-n]; ok && index != i {
			return []int{i, index}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
