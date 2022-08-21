package main

import "fmt"

func generateParenthesis(n int) []string {
	out := []string{}     // output array for results
	loc := make([]int, n) // location of left-parentheses
	for i := 0; i < n; i++ {
		loc[i] = i
	}
	cur := n - 1 // index to left-paren we are moving
	gen(cur, loc, &out)
	return out
}

// recursive function to generate the different permutations
func gen(cur int, loc []int, out *[]string) {
	fmt.Println(loc, cur)
	output(loc, out)
	for (loc[cur] < 2*cur) && (cur != 0) && (cur >= len(loc)-1 || loc[cur]+1 < loc[cur+1]) {
		loc[cur] += 1
		newLoc := make([]int, len(loc))
		copy(newLoc, loc)
		gen(cur-1, newLoc, out)
	}
}

// generate a string of parentheses from location of left-parens
func output(loc []int, out *[]string) {
	s := make([]byte, 2*len(loc))
	j := 0
	for i := 0; i < len(s); i++ {
		if j < len(loc) && i == loc[j] {
			s[i] = '('
			j++
		} else {
			s[i] = ')'
		}
	}
	*out = append(*out, string(s))
}

func main() {
	arr := generateParenthesis(4)
	for i, v := range arr {
		fmt.Println(v, i)
	}
}
