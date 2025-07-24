package main

import "fmt"

func setZeroes(matrix [][]int) {
	mc := [][]int{}
	for i := 0; i < len(matrix); i++ {
		mc = append(mc, matrix[i])
	}
	for i := 0; i < len(mc); i++ {
		for j := 0; j < len(mc[i]); j++ {
			if mc[i][j] != 0 {
				if i>0 &&
			}
		}
	}
	fmt.Println(mc)
}
func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
