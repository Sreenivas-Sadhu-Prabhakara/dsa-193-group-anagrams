package main

import "fmt"

// GroupAnagrams groups words that are anagrams of each other. The grouping
// key is the letter-count array itself — Go arrays are comparable, so a
// [26]int can key a map directly, no sorting and no string building.
func GroupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}
	order := [][26]int{} // first-seen order keeps the output deterministic
	for _, s := range strs {
		var key [26]int
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		if _, ok := groups[key]; !ok {
			order = append(order, key)
		}
		groups[key] = append(groups[key], s)
	}
	out := make([][]string, 0, len(order))
	for _, k := range order {
		out = append(out, groups[k])
	}
	return out
}

func main() {
	fmt.Println(GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
