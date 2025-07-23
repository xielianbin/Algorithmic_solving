package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	idx := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := idx[target-nums[i]]; ok {
			return []int{j, i}
		}
		idx[nums[i]] = i
	}
	return nil
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
