package main

import "fmt"

func solution(a []int, b []int) bool {
	swaps := 0
	lenOfB := len(b)

	for i := range a {
		if a[i] == b[i] {
			continue
		}

		found := false

		for t := i + 1; t < lenOfB; t++ {
			if a[i] == b[t] {
				b[i], b[t] = b[t], b[i]
				found = true
				swaps++
				break
			}
		}

		if !found || swaps > 1 {
			return false
		}
	}

	return true
}

func main() {
	input := [][]int{{9, 8, 3, 4, 5, 6, 7, 2, 1}, {1, 2, 3, 4, 5, 6, 7, 8, 9}}

	fmt.Println("INPUT:", input)
	fmt.Println("OUTPUT:", solution(input[0], input[1]))

}
