package coinchange

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := map[string]struct {
		coins    []int
		amount   int
		expected int
	}{
		"standard coins": {
			[]int{1, 5, 10, 25},
			36,
			3,
		},
		"impossible amount": {
			[]int{2},
			3,
			-1,
		},
		"zero amount": {
			[]int{1},
			0,
			0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := CoinChange(test.coins, test.amount)
			if out != test.expected {
				t.Fatal(fmt.Sprintf("expected %d but got %d", test.expected, out))
			}
		})
	}
}
