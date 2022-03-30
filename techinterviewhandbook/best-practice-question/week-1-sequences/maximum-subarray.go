// Problem: https://leetcode.com/problems/maximum-subarray/

// the solution is we take the first element as the max subarray sum
// we keep aggregating cur_sum as long as it's Positive and update max
// whenever cur_some is negative, we should start looking for a new subarray, keeping the current subarray won't add up the max in any way.

package main

func maxSubArray(nums []int) int {
	max := nums[0]

	curMax := 0

	for _, v := range nums {
		curMax += v

		// keep
		if curMax > max {
			max = curMax
		}

		if curMax < 0 {
			curMax = 0
		}
	}

	return max
}
