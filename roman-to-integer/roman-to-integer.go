package main

import "fmt"

func romanToInt(s string) int {
	out := 0
	i := 0
	last := byte(0)

	for ; i < len(s); i++ {
		switch s[i] {
		case 'M':
			if last == 'C' {
				out -= 200
			}
			out += 1000

		case 'D':
			if last == 'C' {
				out -= 200
			}
			out += 500

		case 'C':
			if last == 'X' {
				out -= 20
			}
			out += 100

		case 'L':
			if last == 'X' {
				out -= 20
			}
			out += 50

		case 'X':
			if last == 'I' {
				out -= 2
			}
			out += 10

		case 'V':
			if last == 'I' {
				out -= 2
			}
			out += 5

		case 'I':
			out++
		}
		last = s[i]
	}
	return out
}

func main() {
	fmt.Println("test")
}
