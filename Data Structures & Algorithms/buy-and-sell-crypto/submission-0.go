func maxProfit(prices []int) int {

	holder := prices[0]
	profit := 0

	for i:=1; i<len(prices); i++ {
		if prices[i] < holder {
			holder = prices[i]
		} else if (prices[i] - holder) > profit {
			profit = prices[i] - holder
		}
	}

	return profit
}