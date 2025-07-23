package main

import (
	"fmt"
	"slices"
)

func findAnagrams(s string, p string) []int {
	bp := []byte(p)
	slices.Sort(bp)
	res := []int{}
	pl := len(p)
	for i := 0; i < len(s)-pl+1; i++ {
		bs := []byte(s[i : i+pl])
		slices.Sort(bs)
		equ := true
		for k := range bs {
			if bs[k] != bp[k] {
				equ = false
				break
			}
		}
		if equ {
			res = append(res, i)
		}
	}
	return res
}
func main() {
	//找到字符串中所有字母异位词
	s := "cbaebabacd"
	p := "abc"
	//s := "abab"
	//p := "ab"

	fmt.Println(findAnagrams(s, p))
}
