package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[l] {
			if nums[l] < target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < nums[r] && nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
func main() {
	nums := []int{3, 4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 4))
}
