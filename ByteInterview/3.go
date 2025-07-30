package main

import "fmt"

func lengthOfLongestSubstring1(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	//定义一个滑动窗口
	bs := []byte(s)
	arr := []byte{}
	maxv := 0
	for _, v := range bs {
		ff := -1
		for i, k := range arr {
			if k == v {
				ff = i
				break
			}
		}
		arr = append(arr[ff+1:], v)
		maxv = max(maxv, len(arr))
	}
	return maxv
}
func lengthOfLongestSubstring2(s string) int {
	//	定义一个滑动窗口
	cnt := [128]int{}
	res := 0
	left := 0
	for right, v := range s {
		cnt[v]++
		for cnt[v] > 1 {
			cnt[s[left]]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

// 无重复的最长字符串
func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring2(s))
}
