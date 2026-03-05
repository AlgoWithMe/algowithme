package twosum

import (
	"fmt"
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := map[string]struct {
		nums     []int
		target   int
		expected []int
	}{
		"first two sum to target": {
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		"last two sum to target": {
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := TwoSum(test.nums, test.target)
			if len(out) != len(test.expected) {
				t.Fatal(fmt.Sprintf("output of TwoSum had len %d but expected has len %d", len(out), len(test.expected)))
			}
			for _, expected := range test.expected {
				if !slices.Contains(out, expected) {
					t.Fatal(fmt.Sprintf("output does not have %d like it should", expected))
				}
			}
		})
	}
}
