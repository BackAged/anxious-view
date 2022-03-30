// Problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

// clarification: problem statement tells we can buy once and sell once on a later day
// not multiple times.

// n*n / bruteForce solution:
//             we buy on ith day and find the max profit by selling on later (i + 1 to n) days
//             we do that for all of the days and update the max

// Observation:
//            the buy day should be always on a earlier day then the sell day

// n Solution:
//        we iterate through the array
//        we update minimum price found so far
//        we calculate the profit against the min price found earlier
//        and update the max profit found so far

package main

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	min := 100005
	max := 0

	for _, v := range prices {
		min = minInt(v, min)
		max = maxInt(max, v-min)
	}

	return max
}
