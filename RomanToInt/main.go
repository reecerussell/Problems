package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Roman To Integer")
	fmt.Println("Roman: MMXIX")
	fmt.Printf("Integer: %d\n", romanToInt("MMXIX"))
}

func romanToInt(s string) int {
	result := 0

	splitString := strings.Split(s, "")

	for i := 0; i < len(s); i++ {
		s1 := getIntFromChar(splitString[i])

		if i+1 < len(s) {
			s2 := getIntFromChar(splitString[i+1])

			if s1 >= s2 {
				result = result + s1
			} else {
				result = result + s2 - s1
				i++
			}
		} else {
			result = result + s1
			i++
		}
	}

	return result
}

func getIntFromChar(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}

	return 0
}
