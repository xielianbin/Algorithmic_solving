package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	s := make([]int, len(nums)+1)
	for i, n := range nums {
		s[i+1] = s[i] + n
	}
	ant := make(map[int]int, len(s))
	res := 0
	for _, sj := range s {
		res += ant[sj-k]
		ant[sj]++
	}
	return res
}
func main() {
	//找到字符串中所有字母异位词
	nums := []int{1, 2, 3}
	k := 3
	//s := "abab"
	//p := "ab"

	fmt.Println(subarraySum(nums, k))
}
