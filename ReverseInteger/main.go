package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Reverse Integer")
	fmt.Println("Input: 36483")
	fmt.Printf("Output: %d\n", reverse(-324879))
}

func reverse(x int) int {

	chars := strings.Split(strconv.Itoa(x), "")
	isNegative := chars[0] == "-"
	if isNegative {
		chars = chars[1:len(chars)]
	}

	reversed := ""
	for i := len(chars) - 1; i >= 0; i-- {
		reversed += chars[i]
	}

	if isNegative {
		reversed = "-" + reversed
	}

	x, _ = strconv.Atoi(reversed)

	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	return x
}
