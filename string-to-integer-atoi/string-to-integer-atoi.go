// https://leetcode.com/problems/string-to-integer-atoi/

package main

import "fmt"

const int32max = 2147483647
const int32min = -2147483648

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func convert(b byte) int {
	return int(b - 48)
}

func myAtoi(s string) int {
	i := 0 // current position
	n := len(s)
	isPositive := true
	result := 0

	// ignore leading whitespace
	for i < n {
		if s[i] != ' ' {
			break
		}
		i++
	}

	// end early if we've reached the end
	if i >= n {
		return 0
	}

	// check for +/- character
	if s[i] == '-' {
		isPositive = false
		i++
	} else if s[i] == '+' {
		i++
	}

	// parse the number-string and build result
	for (i < n) && isNum(s[i]) {
		x := convert(s[i])

		if isPositive {

			if result > (int32max-x)/10 {
				return int32max
			}
			result *= 10
			result += x

		} else {

			if result < (int32min+x)/10 {
				return int32min
			}
			result *= 10
			result -= x

		}
		i++
	}

	return result
}

func main() {
	tests := []string{
		"   +2147483647 blah",
		" -2147483647 asdf",
		" -2147483648 asdf",
		"-3000000000",
		" 2147483648",
		"123.567",
		"    123123123123123123123 asdf ",
		"   adf  123123123123123123123 asdf ",
		"123 123",
		"0000005",
		"000000+5",
		"+0000005",
		"-0000005",
		"   00000000500000",
	}
	for _, s := range tests {
		fmt.Printf("%12d \t\t %#v \n", myAtoi(s), s)
	}
}
