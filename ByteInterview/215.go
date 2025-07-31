package main

import (
	"fmt"
	"math"
)

func findKthLargest(nums []int, k int) int {
	mp := []int{}
	for i := 0; i < k; i++ {
		mp = append(mp, math.MinInt)
	}
	for i := 0; i < len(nums); i++ {
		mpMin := nums[i]
		minIdx := 0
		for j := 0; j < len(mp); j++ {
			if mp[j] < mpMin {
				minIdx = j
				mpMin = mp[j]
			}
		}
		if nums[i] > mpMin {
			mp[minIdx] = nums[i]
		}

	}
	mpMin := math.MaxInt
	for j := 0; j < len(mp); j++ {
		if mp[j] < mpMin {
			mpMin = mp[j]
		}
	}
	return mpMin
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))
}
