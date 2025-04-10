package maxproduct

import (
	"fmt"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := map[string]struct {
		nums     []int
		expected int
	}{
		"all negatives": {
			[]int{-1, -2, -3, -4, -5},
			-6,
		},
		"with a zero": {
			[]int{2, 3, 0, 4, 5},
			60,
		},
		"ascending array": {
			[]int{1, 2, 3, 4, 5},
			60,
		},
		"some negatives": {
			[]int{1, -2, -3, 4, 5},
			30,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := MaxProduct(test.nums)
			if out != test.expected {
				t.Fatal(fmt.Sprintf("expected %d but got %d", test.expected, out))
			}
		})
	}
}
