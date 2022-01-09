package main

import (
	"fmt"
	"math"
)

func solution(n int) int {
	max := 0
	digits := getDigits(n)

	for i := 0; i < len(digits); i++ {
		var d []int

		if i == 0 {
			d = digits[i+1:]
		} else if i == len(digits)-1 {
			d = digits[0:i]
		} else {
			d = append(make([]int, 0), digits[0:i]...)
			d = append(d, digits[i+1:len(digits)]...)
		}

		s := joinDigits(d)

		if s > max {
			max = s
		}
	}

	return max
}

func getDigits(n int) []int {
	digits := make([]int, 1+int(math.Log10(float64(n))))

	i := 0

	for n > 0 {
		d := n % 10
		n /= 10
		digits[len(digits)-i-1] = d
		i++
	}

	return digits
}

func joinDigits(n []int) int {
	t := len(n)

	sum := 0

	for i := 0; i < t; i++ {
		sum += n[i] * int(math.Pow(10, float64(t-i-1)))
	}

	return sum
}

func main() {

	fmt.Println("output:", solution(12))
}
