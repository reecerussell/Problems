package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("String To Integer")
	fmt.Println("Disclaimer: this is my first attempt; it's slow, ugly and uses a fair amount of memory! (9/4/19)")
	fmt.Print("Enter a string: ")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Replace(str, "\n", "", -1)

	fmt.Println("Output: ", myAtoi(str))
}

func myAtoi(str string) int {
	fmt.Println("Input: " + str)

	strs := strings.Split(str, " ")
	for _, e := range strs {

		if e == " " || e == "" {
			continue
		}

		e = regexp.MustCompile("[a-zA-Z]+").Split(e, -1)[0]
		if e == "" {
			return 0
		}

		e = strings.Split(e, ".")[0]
		if e == "" {
			return 0
		}

		chars := strings.Split(e, "")
		e = ""
		for _, c := range chars {
			if len(e) > 0 && (c == "-" || c == "+") {
				break
			}
			e += c
		}

		x := new(big.Int)
		x, ok := x.SetString(e, 10)
		if !ok {
			return 0
		}

		if x.Cmp(big.NewInt(int64(math.MinInt32))) == -1 {
			return math.MinInt32
		}

		if x.Cmp(big.NewInt(int64(math.MaxInt32))) == 1 {
			return math.MaxInt32
		}

		return int(x.Int64())
	}

	return 0
}
