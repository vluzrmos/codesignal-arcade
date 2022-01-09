package main

import (
	"fmt"
	"strings"
)

func solution(inputString string) string {
	start := strings.LastIndex(inputString, "(")

	if start == -1 {
		return inputString
	}

	end := start + 1 + strings.Index(inputString[start+1:], ")")

	sub := helper_StringReverse(helper_StringSub(inputString, start+1, end))

	newInput := helper_StringSub(inputString, 0, start) + sub + helper_StringSub(inputString, end+1, len(inputString))

	return solution(newInput)
}

func helper_StringReverse(input string) string {
	if len(input) <= 1 {
		return input
	}

	s := []rune(input)

	for i, j := 0, len(s)-1; i <= (len(s)-1)/2; i++ {
		s[i], s[j-i] = s[j-i], s[i]
	}

	return string(s)
}

func helper_StringSub(input string, start, end int) string {
	if start >= len(input) || end > len(input) {
		return ""
	}

	if start < 0 || end <= 0 {
		return ""
	}

	if start >= end {
		return ""
	}

	return input[start:end]
}

func main() {
	inputs := [...]string{
		"(bar)",
		"foo(bar)baz",
		"foo(bar)baz(blim)",
		"foo(bar(baz))blim",
		"()",
	}

	expectations := [...]string{
		"rab",
		"foorabbaz",
		"foorabbazmilb",
		"foobazrabblim",
		"",
	}

	for i, inputString := range inputs {
		output := solution(inputString)
		expected := expectations[i]

		fmt.Println("INPUT:", inputString)
		fmt.Println("OUTPUT:", output)
		fmt.Println("EXPECTED:", expected)
		fmt.Println("IS CORRECT? ", expected == output)
		fmt.Println()
	}

}
