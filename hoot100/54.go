package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	maxrow := len(matrix) - 1
	maxcol := len(matrix[0]) - 1
	minrow := 0
	mincol := 0
	r := 0
	c := 0
	res := []int{}
	for minrow <= maxrow && mincol <= maxcol {
		if r == minrow && minrow <= maxrow && mincol <= maxcol {
			for c <= maxcol {
				res = append(res, matrix[r][c])
				c++
			}
			c--
			r++
			minrow++
		}
		if c == maxcol && minrow <= maxrow && mincol <= maxcol {
			for r <= maxrow {
				res = append(res, matrix[r][c])
				r++
			}
			r--
			c--
			maxcol--
		}
		if r == maxrow && minrow <= maxrow && mincol <= maxcol {
			for c >= mincol {
				res = append(res, matrix[r][c])
				c--
			}
			c++
			r--
			maxrow--
		}
		if c == mincol && minrow <= maxrow && mincol <= maxcol {
			for r >= minrow {
				res = append(res, matrix[r][c])
				r--
			}
			r++
			c++
			mincol++
		}
	}
	return res
}
func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	fmt.Println(spiralOrder(matrix))
}
