package besttimetobuyandsellstock

// MaxProfit returns the maximum profit from buying and selling a stock once.
func MaxProfit(prices []int) int {
	max, profit, left, right := 0, 0, 0, 1
	for right < len(prices) {
		if prices[left] <= prices[right] {
			if profit < prices[right] {
				profit = prices[right]
			}

			right++
		} else if prices[left] > prices[right] {
			profit = 0
			left = right
			right++
		}

		if profit > max {
			max = profit - prices[left]
		}
	}

	return max
}
