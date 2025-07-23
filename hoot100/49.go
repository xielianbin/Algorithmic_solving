package main

import (
	"fmt"
	"maps"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	res := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		bs := []byte(strs[i])
		slices.Sort(bs)
		tmp := string(bs)
		res[tmp] = append(res[tmp], strs[i])
	}
	ans := [][]string{}
	for _, m := range res {
		ans = append(ans, m)
	}
	tt := maps.Values(res)
	ss := slices.Collect(tt)
	fmt.Println(ss)
	return ans
}
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
}
