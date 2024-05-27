// https://leetcode.com/problems/best-time-to-buy-and-sell-stock
package main

import "fmt"

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(nums))
}

func maxProfit(prices []int) int {

	min, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		cost := prices[i] - min
		if profit < cost {
			profit = cost
		}
		if min > prices[i] {
			min = prices[i]
		}
	}

	return profit
}
