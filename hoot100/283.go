package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	count := 0
	if len(nums) < 2 && nums[0] == 0 {
		return
	}
	for _, n := range nums {
		if n == 0 {
			count++
		}
	}
	m := 0
	for i := 0; i < len(nums); i++ {

		if i+m > len(nums)-1 {
			nums[i] = 0
		} else {
			j := i + m
			for ; j < len(nums)-1; j++ {
				if nums[j] == 0 {
					m++
				} else {
					break
				}
			}
			nums[i] = nums[i+m]
		}
	}
}
func main() {
	//nums := []int{0, 1, 0, 3, 12}
	nums := []int{0}
	//nums := []int{100, 4, 200, 1, 3, 2}
	moveZeroes(nums)
	fmt.Println(nums)
}
