// https://leetcode.com/problems/letter-combinations-of-a-phone-number/submissions/

package main

import (
	"fmt"
)

var numMap = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {

	// end early if provided with no input
	if len(digits) == 0 {
		return []string{}
	}

	// allocate result array
	resultSize := 1
	for _, d := range digits {
		resultSize *= len(numMap[d])
	}
	result := make([]string, resultSize)

	// allocate pointer table (this defines the current permutation)
	pointers := make([]int, len(digits))
	for i := range pointers {
		pointers[i] = 0
	}

	// buffer to hold each individual result
	b := make([]byte, len(digits))

	// fill each empty spot in the result array
	for k := 0; ; k++ {

		// save the current permutation to result array
		for i, digit := range digits {
			x := numMap[digit]
			y := pointers[i]
			b[i] = x[y]
		}
		result[k] = string(b)

		// increment the permutation
		for j := len(pointers) - 1; j >= 0; j-- {
			key := rune(digits[j])
			pointers[j] = (pointers[j] + 1) % len(numMap[key])
			if pointers[j] != 0 {
				break
			}
			if j == 0 {
				goto end
			}
		}
	}
end:
	return result
}

func main() {
	inputs := []string{
		"237",
	}
	for _, v := range inputs {
		fmt.Println(letterCombinations(v))
	}
}
