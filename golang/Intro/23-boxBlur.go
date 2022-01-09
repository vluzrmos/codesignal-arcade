package main

import "fmt"

func solution(image [][]int) [][]int {
	result := make([][]int, 0)

	for i := 0; i+3 <= len(image); i++ {
		row := make([]int, 0)

		for j := 0; j+3 <= len(image[0]); j++ {
			row = append(row, MatrixSumOffset(image, i, j)/9)
		}

		result = append(result, row)
	}

	return result
}

func MatrixSumOffset(matrix [][]int, iOffset, jOffset int) int {
	sum := 0

	for i := iOffset; i < 3+iOffset; i++ {
		for j := jOffset; j < 3+jOffset; j++ {
			sum += matrix[i][j]
		}
	}

	return sum
}

func main() {

	input := [][]int{
		{36, 0, 18, 9, 9, 45, 27},
		{27, 0, 54, 9, 0, 63, 90},
		{81, 63, 72, 45, 18, 27, 0},
		{0, 0, 9, 81, 27, 18, 45},
		{45, 45, 27, 27, 90, 81, 72},
		{45, 18, 9, 0, 9, 18, 45},
		{27, 81, 36, 63, 63, 72, 81},
	}

	expected := "[[39 30 26 25 31] [34 37 35 32 32] [38 41 44 46 42] [22 24 31 39 45] [37 34 36 47 59]]"

	fmt.Println("Solution:")
	fmt.Println(solution(input))

	fmt.Println("Expected:")
	fmt.Println(expected)
}
