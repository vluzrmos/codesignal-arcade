package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []string{
		"f",
		"g",
		"a",
		"h",
	}

	expected := true

	output := solution(input)

	fmt.Println("input:", input)
	fmt.Println("output:", output)
	fmt.Println("expected:", expected)

	if output != expected {
		fmt.Println("!!!Incorreto!!!")
	}
}

func solution(inputArray []string) bool {
	wg := sync.WaitGroup{}

	wg.Add(len(inputArray))

	result := make(chan bool, len(inputArray))

	for i := 0; i < len(inputArray); i++ {
		go func(i int, wg *sync.WaitGroup, result chan bool) {
			defer func(wg *sync.WaitGroup) {
				wg.Done()
				fmt.Println("Ending:", i)
			}(wg)

			fmt.Println("Starting:", i)

			result <- findPath(inputArray, i, []int{i}, 1)
		}(i, &wg, result)
	}

	go func(wg *sync.WaitGroup, result chan bool) {
		defer close(result)
		wg.Wait()
	}(&wg, result)

	for r := range result {
		if r {
			return true
		}
	}

	return false
}

func findPath(inputArray []string, current int, except []int, n int) bool {
	if n == len(inputArray) {
		return true
	}

	possibilities := findAllPossibilities(inputArray, current, except)

	if len(possibilities) == 0 {
		return false
	}

	found := false

	for _, i := range possibilities {
		if i == current || sliceIntsContains(except, i) {
			continue
		}

		ignoring := append(make([]int, 0), except...)

		ignoring = append(ignoring, i)

		found = findPath(
			inputArray,
			i,
			ignoring,
			n+1,
		)

		if found {
			return true
		}
	}

	return found
}

func findAllPossibilities(input []string, current int, except []int) []int {
	keys := make([]int, 0)

	for i := 0; i < len(input); i++ {
		if (i == current) || sliceIntsContains(except, i) {
			continue
		}

		if StrDiffersOneChar(input[current], input[i]) {
			keys = append(keys, i)
		}
	}

	return keys
}

func StrDiffersOneChar(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	count := 0

	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return count == 1
}

func sliceIntsContains(haystack []int, needle int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}
