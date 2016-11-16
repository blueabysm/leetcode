package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var begin int = 0
	var max_length int = 0
	var current_length int = 0
	var occur_map map[string]int
	occur_map = make(map[string]int)

	for i, c := range s {
		if index, ok := occur_map[string(c)]; ok {
			if index >= begin {
				begin = index + 1
			}
		}
		occur_map[string(c)] = i
		current_length = i - begin + 1
		if current_length > max_length {
			max_length = current_length
			fmt.Println("current max s[", begin, ":", i, "] = ", s[begin:i])
		}
	}
	return max_length
}

func main() {
	str := "211n;no30q3n4q23402noqto2n43"
	fmt.Println(str)
	fmt.Println(lengthOfLongestSubstring(str))
}
