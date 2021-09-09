package leetcode

import "fmt"

const (
	topic = "Leetcode Problem 0121. 买卖股票的最佳时机\n"
)

func init() {
	fmt.Println(topic)
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	i := 0
	j := 0
	pointer := i + 1

	maxProfit := prices[j] - prices[i]

	// [7,1,5,3,6,4]
	for pointer < len(prices) {
		// 找到局部最小值，就是 buy point
		if prices[pointer] < prices[i] {
			// update buy point
			i = pointer
		}
		profit := prices[pointer] - prices[i]
		if profit > maxProfit {
			// update sell point
			j = pointer
			maxProfit = profit
		}
		pointer++
	}
	return maxProfit
}

func bruteForce(prices []int) int {
	maxProfit := 0
	for i, buy := range prices {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - buy
			if maxProfit < profit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
