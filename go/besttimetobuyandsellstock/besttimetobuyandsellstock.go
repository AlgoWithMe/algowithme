package besttimetobuyandsellstock

func MaxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for _, p := range prices[1:] {
		maxProfit = max(maxProfit, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return maxProfit
}
