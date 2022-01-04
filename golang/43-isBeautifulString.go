package main

import (
	"fmt"
	"strings"
)

func solution(inputString string) bool {
	counts := makeLettersMap(inputString)

	for i := 1; i < len(counts); i++ {
		if counts[i] > counts[i-1] {
			return false
		}
	}

	return true
}

func makeLettersMap(inputString string) []int {
	counts := make([]int, 0)
	chars := strings.Split(inputString, "")

	uniq := make(map[string]int, 0)

	index := 0

	for i := 'a'; i <= 'z'; i++ {
		uniq[string(i)] = index
		counts = append(counts, 0)
		index++
	}

	for _, c := range chars {
		if key, found := uniq[c]; found {
			counts[key]++
			continue
		}
	}

	return counts
}

func main() {
	fmt.Println(solution("abcdefghijklmnopqrstuvwxyzz"))
}
