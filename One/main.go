package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{10, 15, 3, 7}
	k := 17

	fmt.Println("Programming Problem #1")
	fmt.Println("List: [10, 15, 3, 7]")
	fmt.Println("K: 17")

	result := canSumToK(arr, k)

	fmt.Println("Result: " + strconv.FormatBool(result))
}

func canSumToK(arr []int, k int) (result bool) {

	for _, i := range arr {
		for _, j := range arr {
			if (i+j) == k && i != j {
				fmt.Printf("%d + %d = %d\n", i, j, k)
				return true
			}
		}
	}

	return false
}
