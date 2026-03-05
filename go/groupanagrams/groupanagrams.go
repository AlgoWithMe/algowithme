package groupanagrams

import "slices"

func GroupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, s := range strs {
		key := []byte(s)
		slices.Sort(key)
		groups[string(key)] = append(groups[string(key)], s)
	}
	result := make([][]string, 0, len(groups))
	for _, g := range groups {
		result = append(result, g)
	}
	return result
}
