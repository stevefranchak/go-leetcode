package buy_sell_stock

func maxProfit(prices []int) int {
	maxProfit := 0
	currentMinPrice := 10001 // one beyond max possible price

	for _, p := range prices {
		if p < currentMinPrice {
			currentMinPrice = p
			continue
		}
		potentialProfit := p - currentMinPrice
		if potentialProfit > maxProfit {
			maxProfit = potentialProfit
		}
	}

	return maxProfit
}
