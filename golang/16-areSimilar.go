package main

import (
	"fmt"
)

func solution(a []int, b []int) bool {
	c := make([]int, 0)
	d := make([]int, 0)

	equals := true

	for i := range a {
		if a[i] != b[i] {
			equals = false
			c = append(c, a[i])
			d = append(d, b[i])
		}
	}

	if equals {
		return true
	}

	if len(c) != 2 {
		return false
	}

	d[0], d[1] = d[1], d[0]

	for i := range c {
		if c[i] != d[i] {
			return false
		}
	}

	return true
}

func main() {
	input := [][]int{{1, 2, 3, 4, 5, 6, 7, 9, 8}, {1, 2, 3, 4, 5, 6, 7, 8, 9}}

	fmt.Println("INPUT:", input)
	fmt.Println("OUTPUT:", solution(input[0], input[1]))

}
