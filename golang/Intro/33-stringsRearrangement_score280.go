package main

import (
	"sort"
)

func solution(inputArray []string) bool {
	var output = append([]string{}, inputArray...)

	sort.Strings(output)

	if IsRearrangeable(inputArray) || IsRearrangeable(output) {
		return true
	}

	keys := make(map[int]int)

	for i := 0; i < len(inputArray); i++ {
		output = append([]string{}, inputArray[i])

		current := i
		keys[i] = i

		for true {
			next, ok := findNextWithDiffersOne(inputArray, current, keys)

			if !ok {
				keys = make(map[int]int)
				break
			}

			current = next
			keys[next] = next
			output = append(output, inputArray[next])
		}

		if len(output) == len(inputArray) && IsRearrangeable(output) {
			return true
		}
	}

	return false
}

func findNextWithDiffersOne(input []string, current int, except map[int]int) (int, bool) {
	for i := 0; i < len(input); i++ {
		if _, found := except[i]; found {
			continue
		}

		if StrDiffersOneChar(input[current], input[i]) {
			return i, true
		}
	}

	return -1, false
}

func IsRearrangeable(str []string) bool {
	for i := 0; i < len(str)-1; i++ {
		if !StrDiffersOneChar(str[i], str[i+1]) {
			return false
		}
	}

	return true
}

func StrDiffersOneChar(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	count := 0

	for i, _ := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return count == 1
}
