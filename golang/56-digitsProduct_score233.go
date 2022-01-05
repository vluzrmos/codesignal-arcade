package main

import (
	"fmt"
	"math"
	"sort"
)

func solution(product int) int {
	if product == 0 {
		return 10
	}

	if product < 10 {
		return product
	}

	if d, found := baseOfTen(product); found {
		return d
	}

	if isPrime(product) {
		return -1
	}

	divs := IntDivs(product)

	for _, div := range divs {
		if div > 9 {
			return -1
		}
	}

	return JoinInts(CombineInts(divs))
}

func isPrime(n int) bool {
	if n < 1 {
		return false
	}

	if n == 2 {
		return true
	}

	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func baseOfTen(n int) (int, bool) {
	exp := int(math.Pow(10, float64(int(math.Log10(float64(n))))))

	if n%exp == 0 {
		return n / exp, true
	}

	return n, false
}

func IntDivs(n int) []int {
	divs := make([]int, 0)

	i := 2

	for n > 1 {
		if n%i != 0 {
			i++
			continue
		}

		divs = append(divs, i)
		n = n / i
	}

	return divs
}

func CombineInts(n []int) []int {
	result := make([]int, 0)

	for i := len(n) - 1; i >= 0; i-- {
		r := n[i]

		for j := i - 1; j >= 0; j-- {
			if r*n[j] > 9 {
				break
			}
			i = j
			r *= n[j]
		}

		result = append(result, r)
	}

	sort.Ints(result)

	return result
}

func JoinInts(n []int) int {
	result := 0
	exp := int(math.Pow(10, float64(len(n)-1)))

	for i := 0; i < len(n); i++ {
		result += n[i] * exp
		exp /= 10
	}

	return result
}

type Test struct {
	Input    int
	Expected int
}

func main() {
	tests := []Test{
		{
			Input:    12,
			Expected: 26,
		},
		{
			Input:    10,
			Expected: 1,
		},
		{
			Input:    100,
			Expected: 1,
		},
		{
			Input:    1000,
			Expected: 1,
		},
		{
			Input:    10000,
			Expected: 1,
		},
		{
			Input:    600,
			Expected: 6,
		},
	}

	count := 0

	for index, test := range tests {
		fmt.Println()

		fmt.Println("### Test", index, "###")

		output := solution(test.Input)
		correct := test.Expected == output

		fmt.Println("Input:", test.Input)
		fmt.Println("Output:", output)
		fmt.Println("Expected:", test.Expected)
		fmt.Println("Correct:", correct)
		if correct {
			count++
		}
	}

	fmt.Printf("%d/%d\n", count, len(tests))
}
