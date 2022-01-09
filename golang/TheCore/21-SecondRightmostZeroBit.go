package main

import "fmt"

// Not the solution yet
func solution(n int) int {
	fmt.Printf("%b\n", ^uint(n))

	return int(^uint(0) &^ uint(n))
}

func main() {
	solution(37)
}
