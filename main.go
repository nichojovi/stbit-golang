package main

import (
	"fmt"
	"sort"
)

func sortCharacter(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	mm := make(map[string][]string)
	for _, v := range strs {
		str := sortCharacter(v)
		mm[str] = append(mm[str], v)
	}

	var res [][]string
	for _, v := range mm {
		res = append(res, v)
	}
	return res
}

func main() {
	var ss = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	fmt.Println(groupAnagrams(ss))
}
