package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("Hexadecimal To Decimal")
	fmt.Println("Input: 7E3")
	fmt.Println("Expected: 2019")
	fmt.Printf("Output: %d\n", Decode("7E3"))
}

func Decode(hex string) int {
	var output int = 0

	hexMap := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"A": 10,
		"B": 11,
		"C": 12,
		"D": 13,
		"E": 14,
		"F": 15,
	}

	chars := strings.Split(hex, "")

	for i, e := range chars {
		power := int(math.Pow(16, float64(len(chars)-1-i)))
		output += (hexMap[e] * power)
	}

	return output
}
