package main

import (
	"fmt"
	"slices"
)

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < len(matrix); i++ {
		slices.Reverse(matrix[i])
	}
}
func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	rotate(matrix)
	fmt.Println(matrix)
}
