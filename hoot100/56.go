package main

import (
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(p, q []int) int {
		return p[0] - q[0]
	})
	res := [][]int{}
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			if res[len(res)-1][1] < intervals[i][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
func main() {
	intervals := [][]int{{1, 3}, {15, 18}, {2, 6}, {8, 10}}
	//intervals := [][]int{{1, 4}, {0, 4}}
	fmt.Println(merge(intervals))
}
