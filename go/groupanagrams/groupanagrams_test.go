package groupanagrams

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func normalize(groups [][]string) []string {
	result := make([]string, len(groups))
	for i, group := range groups {
		sorted := make([]string, len(group))
		copy(sorted, group)
		sort.Strings(sorted)
		key := fmt.Sprintf("%v", sorted)
		result[i] = key
	}
	sort.Strings(result)
	return result
}

func TestGroupAnagrams(t *testing.T) {
	tests := map[string]struct {
		strs     []string
		expected [][]string
	}{
		"mixed anagram groups": {
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		"single empty string": {
			[]string{""},
			[][]string{{""}},
		},
		"single character": {
			[]string{"a"},
			[][]string{{"a"}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := GroupAnagrams(test.strs)
			if len(out) != len(test.expected) {
				t.Fatal(fmt.Sprintf("expected %d groups but got %d", len(test.expected), len(out)))
			}
			normalizedOut := normalize(out)
			normalizedExpected := normalize(test.expected)
			if !slices.Equal(normalizedOut, normalizedExpected) {
				t.Fatal(fmt.Sprintf("expected %v but got %v", normalizedExpected, normalizedOut))
			}
		})
	}
}
