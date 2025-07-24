package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	res := math.MinInt
	f := 0
	for i := 0; i < len(nums); i++ {
		f = max(f, 0) + nums[i]
		res = max(res, f)
	}

	return res
}
func main() {
	nums := []int{-2, -1}
	fmt.Println(maxSubArray(nums))
}
