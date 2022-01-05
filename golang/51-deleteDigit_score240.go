import "math"

func solution(n int) int {
	max := 0
	digits := int(math.Log10(float64(n)) + 1)

	if digits == 2 {
		if n/10 > n%10 {
			return n / 10
		}

		return n % 10
	}

	for i := 0; i < digits; i++ {
		left := n / int(math.Pow(10, float64(digits-i)))
		right := int(n % int(math.Pow(10, float64(digits-i-1))))

		t := int(math.Pow(10, float64(int(math.Log10(float64(right))+1))))

		amount := left*t + right

		if amount > max {
			max = amount
		}
	}

	return max
}

func minDigit(n int) int {
	min := n % 10
	n /= 10

	for n > 0 {
		d := n % 10
		n /= 10

		if d < min {
			min = d
		}
	}

	return min
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
