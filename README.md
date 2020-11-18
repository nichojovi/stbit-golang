##### Stbit Project
```
Brought to you by Nicholas J.
Programming language: Go
Database: MySQL
How to run: make run
```

##### 1. User Query
```
select
	A.id,
	A.userName,
	B.userName ParentUserName
FROM
	user A
LEFT JOIN 
	user B
ON A.Parent = B.id
```

##### 2. Omdbapi Project
```
On this repository
1. API was cover with authentication by username and password (standard encrypted SHA1)
2. Get data from ombdapi
3. Store log on DB (mysql), table provided
4. Using go routine in service level
5. Sample unit test on repository level
6. Clean code architecture
```

##### 3. Refactor code
```
func findFirstStringInBracket(str string) string {
	if (len(str) > 0) {
		indexFirstBracketFound := strings.Index(str,"(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket,")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound-1])
			}
		}
	}
	return ""
}
```

##### 4. Anagram code
```
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
```