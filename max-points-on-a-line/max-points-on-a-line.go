package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type lineMap struct {
	m map[[2]string]map[int]bool
}

func (l *lineMap) add(entry [2]string, pointIndex int) {
	if _, ok := l.m[entry]; !ok {
		l.m[entry] = make(map[int]bool)
	}
	l.m[entry][pointIndex] = true
}

func _maxPoints(points [][]int) int {
	l := lineMap{m: make(map[[2]string]map[int]bool)}
	best := 0
	n := len(points)

	if len(points) == 1 {
		return 1
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a := points[i]
			b := points[j]

			// vertical lines are special, and are handled seperately
			if a[0] == b[0] {
				key := [2]string{"vert", strconv.Itoa(a[0])}
				l.add(key, i)
				l.add(key, j)
				continue
			}

			slope := big.NewRat(int64(b[1]-a[1]), int64(b[0]-a[0]))

			intercept := new(big.Rat)
			x := big.NewRat(int64(a[0]), 1)
			y := big.NewRat(int64(a[1]), 1)
			z := new(big.Rat)
			z.Mul(slope, x)
			intercept.Sub(y, z)

			key := [2]string{slope.String(), intercept.String()}
			l.add(key, i)
			l.add(key, j)
		}
	}

	for _, v := range l.m {
		if x := len(v); x > best {
			best = x
		}
	}
	return best
}

func (l *lineMap) display() {
	println("map:")
	for k, v := range l.m {
		fmt.Printf("%8.4v: %v \t\t", k, len(v))
		for x := range v {
			fmt.Printf("%v,", x)
		}
		fmt.Println()
	}
}

// =============================================================================
// This is a simpler solution

func maxPoints(points [][]int) int {
	slope := 0.0
	best := 1
	for i := 0; i < len(points); i++ {
		m := map[float64]int{}
		ax := points[i][0]
		ay := points[i][1]

		for j := i + 1; j < len(points); j++ {
			bx := points[j][0]
			by := points[j][1]

			if ax == bx {
				slope = math.MaxFloat64
			} else {
				slope = float64(by-ay) / float64(bx-ax)
			}

			if _, ok := m[slope]; !ok {
				m[slope] = 2
			} else {
				m[slope]++
			}

			if m[slope] > best {
				best = m[slope]
			}
		}
		display(m)
	}
	return best
}

func display(m map[float64]int) {
	for k, v := range m {
		fmt.Printf("%8.4v: \t %v", k, v)
		fmt.Println()
	}
}

// =============================================================================

func main() {
	// points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	// points := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	// points := [][]int{{0, 0}}

	// TODO: expected result = 6
	points := [][]int{{7, 3}, {19, 19}, {-16, 3}, {13, 17}, {-18, 1}, {-18, -17}, {13, -3}, {3, 7}, {-11, 12}, {7, 19}, {19, -12}, {20, -18}, {-16, -15}, {-10, -15}, {-16, -18}, {-14, -1}, {18, 10}, {-13, 8}, {7, -5}, {-4, -9}, {-11, 2}, {-9, -9}, {-5, -16}, {10, 14}, {-3, 4}, {1, -20}, {2, 16}, {0, 14}, {-14, 5}, {15, -11}, {3, 11}, {11, -10}, {-1, -7}, {16, 7}, {1, -11}, {-8, -3}, {1, -6}, {19, 7}, {3, 6}, {-1, -2}, {7, -3}, {-6, -8}, {7, 1}, {-15, 12}, {-17, 9}, {19, -9}, {1, 0}, {9, -10}, {6, 20}, {-12, -4}, {-16, -17}, {14, 3}, {0, -1}, {-18, 9}, {-15, 15}, {-3, -15}, {-5, 20}, {15, -14}, {9, -17}, {10, -14}, {-7, -11}, {14, 9}, {1, -1}, {15, 12}, {-5, -1}, {-17, -5}, {15, -2}, {-12, 11}, {19, -18}, {8, 7}, {-5, -3}, {-17, -1}, {-18, 13}, {15, -3}, {4, 18}, {-14, -15}, {15, 8}, {-18, -12}, {-15, 19}, {-9, 16}, {-9, 14}, {-12, -14}, {-2, -20}, {-3, -13}, {10, -7}, {-2, -10}, {9, 10}, {-1, 7}, {-17, -6}, {-15, 20}, {5, -17}, {6, -6}, {-11, -8}}

	result := maxPoints(points)
	fmt.Println("FINAL RESULT =", result)

}
