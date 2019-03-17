package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 0, -2}

	fmt.Println("Programming Problem #2")
	fmt.Println("Input: [9,0,-2]")
	fmt.Println("Expected: [0,-18,0]")

	arr = products(arr)

	fmt.Printf("Output: [%d, %d, %d]\n", arr[0], arr[1], arr[2])
}

func products(input []int) []int {
	length := len(input)
	output := []int{}

	sum := 1
	index := 0
	for i := 0; i < length; i++ {

		if i != index {
			sum *= input[i]
		}

		if i == length-1 {
			output = append(output, sum)

			if len(output) < length {
				sum = 1
				i = 0
				index++
			}
		}
	}

	fmt.Println(input)
	fmt.Println(output)

	return output
}
