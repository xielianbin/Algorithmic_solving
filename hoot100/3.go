package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	maxn := 0
	if len(s) < 2 {
		return len(s)
	}
	bs := []byte(s)
	arr := []byte{}

	for _, v := range bs {
		ff := false
		for _, a := range arr {
			if a == v {
				ff = true
			}
		}
		count := 0
		for n := range arr {
			if arr[n] {
				count++
			}
		}
		maxn = max(maxn, count)
	}
	return maxn
}
func main() {
	//nums := []int{0, 1, 0, 3, 12}
	//nums := "abcabcbb"
	//nums := "au"
	nums := "aab"
	//nums := "abba"
	//nums := "pwwkew"
	//nums := []int{100, 4, 200, 1, 3, 2}

	fmt.Println(lengthOfLongestSubstring(nums))
}
