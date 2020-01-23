/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit. You cannot sell a stock before you buy one.
*/

package leetcode

import "math"

// One Pass solution O(n)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min := math.MaxInt32
	max := 0

	for i, p := range prices {
		if p < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}

	return max
}

// Bruteforce solution O(n^2)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	output := 0
	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			sum := prices[i] - prices[j]
			if sum > output {
				output = sum
			}
		}
	}
	return output
}
