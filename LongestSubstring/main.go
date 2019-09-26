package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "babad"
	fmt.Printf("Input: %s\n", input)

	fmt.Printf("%d", lengthOfLongestSubstring(input))
}
func lengthOfLongestSubstring(s string) int {
	inputArr := strings.Split(s, "")
	arr := []string{}
	for i := 0; i < len(inputArr); i++ {
		cs := ""
		for j := i; j < len(inputArr); j++ {
			c := strings.Contains(cs, inputArr[j])
			if !c {
				cs += inputArr[j]
			} else {
				break
			}
		}
		arr = append(arr, cs)
	}

	b := 0
	for _, e := range arr {
		l := len(e)
		if l > b {
			b = l
		}
	}

	return b
}
