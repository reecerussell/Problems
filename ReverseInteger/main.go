package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Reverse Integer")
	fmt.Println("Input: 36483")
	fmt.Printf("Output: %d\n", reverse(36483))
}

func reverse(x int) int {

	num := strconv.Itoa(x)

	isNegative := strings.Index(num, "-") > -1

	if isNegative {
		num = strings.Replace(num, "-", "", 1)
	}

	reversedNum := ""

	chars := strings.Split(num, "")

	for i := len(chars) - 1; i >= 0; i-- {
		reversedNum += chars[i]
	}

	if isNegative {
		num = "-" + reversedNum
	} else {
		num = reversedNum
	}

	y, err := strconv.ParseInt(num, 10, 32)
	if err != nil {
		return 0
	}

	x = int(y)

	return x
}
