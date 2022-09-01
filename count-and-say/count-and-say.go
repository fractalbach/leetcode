package main

import "fmt"

var memo = map[int]string{}

func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}
	if val, ok := memo[n]; ok {
		return val
	}
	s := countAndSay(n - 1)
	output := ""

	var digit byte = s[0]
	var count int = 0

	for i := 0; i < len(s); i++ {
		if s[i] != digit {
			output += fmt.Sprintf("%v%v", count, string(digit))
			digit = s[i]
			count = 1
		} else {
			count++
		}
	}
	output += fmt.Sprintf("%v%v", count, string(digit))

	memo[n] = output
	return output
}

func main() {
	for i := 1; i < 11; i++ {
		fmt.Println(i, countAndSay(i))
	}
}
