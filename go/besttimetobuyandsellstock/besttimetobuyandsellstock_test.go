package besttimetobuyandsellstock

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := map[string]struct {
		prices   []int
		expected int
	}{
		"buy low sell high": {
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		"prices only decrease": {
			[]int{7, 6, 4, 3, 1},
			0,
		},
		"two elements profit": {
			[]int{1, 2},
			1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := MaxProfit(test.prices)
			if out != test.expected {
				t.Fatal(fmt.Sprintf("expected %d but got %d", test.expected, out))
			}
		})
	}
}
