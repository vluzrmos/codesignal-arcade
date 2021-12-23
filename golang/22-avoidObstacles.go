package main

import (
	"fmt"
	"sort"
)

func solution(inputArray []int) int {
	sort.Ints(inputArray)

	max := inputArray[len(inputArray)-1]
	jump := 1

	indexes := make(map[int]int, 0)

	for _, v := range inputArray {
		indexes[v] = v
	}

	for i := jump; i <= max+1; i += jump {
		if _, found := indexes[i]; !found {
			continue
		}

		jump++
		i = 0
	}

	return jump
}

func main() {
	input := []int{1000, 999}

	fmt.Println(solution(input))
}
