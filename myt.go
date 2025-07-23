package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	newIntervals := make([][]int, 0)
	newIntervals = append(newIntervals, intervals[0])
	for i := 0; i < len(intervals); i++ {
		m := len(newIntervals) - 1
		if intervals[i][0] <= newIntervals[m][1] {
			if newIntervals[m][1] < intervals[i][1] {
				newIntervals[m][1] = intervals[i][1]
			}

		} else {
			newIntervals = append(newIntervals, intervals[i])
		}
	}
	return newIntervals
}
func main() {
	intervals := [][]int{
		{1, 3}, {2, 6}, {15, 18}, {8, 10},
	}
	intervals = merge(intervals)
	fmt.Println(intervals)
}
