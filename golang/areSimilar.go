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

		for t := i + 1; swaps == 0 && t < lenOfB; t++ {
			if a[i] == b[t] {
				b[i], b[t] = b[t], b[i]
				found = true
				swaps++
				break
			}
		}

		if !found {
			return false
		}
	}

	return swaps <= 1
}

func main() {
	input := [][]int{{1, 2, 3, 4, 5, 6, 7, 9, 8}, {1, 2, 3, 4, 5, 6, 7, 8, 9}}

	fmt.Println("INPUT:", input)
	fmt.Println("OUTPUT:", solution(input[0], input[1]))

}
