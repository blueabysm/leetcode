package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	if len(numbers) <= 0 {
		return nil
	}

	var left, right int = 0, len(numbers) - 1

	for {
		if left >= right {
			return nil
		}

		if numbers[left]+numbers[right] < target {
			left++
			continue
		}
		if numbers[left]+numbers[right] > target {
			right--
			continue
		}
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
	}
}

func main() {
	fmt.Println(twoSum([]int{-10, -4, -2, -1, 3, 4, 5, 6, 7}, 0))
}
